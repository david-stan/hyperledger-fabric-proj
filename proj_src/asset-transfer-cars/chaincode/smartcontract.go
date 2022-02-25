package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// Asset describes basic details of what makes up a simple asset
//Insert struct field in alphabetic order => to achieve determinism accross languages
// golang keeps the order when marshal to json but doesn't order automatically
type MalfunctionAsset struct {
	Desc  string  `json:"Description"`
	Price float64 `json:"Price"`
}

type CarAsset struct {
	ID              string             `json:"ID"`
	Brand           string             `json:"Brand"`
	Model           string             `json:"Model"`
	Year            string             `json:"Year"`
	Colour          string             `json:"Colour"`
	Owner_Id        string             `json:"Owner_Id"`
	MalfunctionList []MalfunctionAsset `json:"MalfunctionList"`
}

type PersonAsset struct {
	ID      string  `json:"ID"`
	Name    string  `json:"Name"`
	Surname string  `json:"Surname"`
	Email   string  `json:"Email"`
	Credit  float64 `json:"Credit"`
}

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {

	personAssetList := []PersonAsset{
		{
			ID:      "p_0001",
			Name:    "Milomir",
			Surname: "Maric",
			Email:   "milomirmaric@gmail.com",
			Credit:  1000.0,
		},
		{
			ID:      "p_0002",
			Name:    "Mihajlo",
			Surname: "Ulemek",
			Email:   "mihajloulemek@gmail.com",
			Credit:  500.0,
		},
		{
			ID:      "p_0003",
			Name:    "Zeljko",
			Surname: "Raznjatovic",
			Email:   "zeljkozeljko@gmail.com",
			Credit:  20000.0,
		},
	}

	carAssetList := []CarAsset{
		{
			ID:       "c_0001",
			Brand:    "Mazda",
			Model:    "CX-5",
			Year:     "2015",
			Colour:   "Cherry Red",
			Owner_Id: "0001",
			MalfunctionList: []MalfunctionAsset{
				{
					Desc:  "Breaks",
					Price: 100.0,
				},
				{
					Desc:  "MAF Sensor",
					Price: 600.0,
				},
			},
		},
		{
			ID:       "c_0002",
			Brand:    "Toyota",
			Model:    "Corolla",
			Year:     "2009",
			Colour:   "Mettalic Gray",
			Owner_Id: "0001",
			MalfunctionList: []MalfunctionAsset{
				{
					Desc:  "Loose timing belt",
					Price: 100.0,
				},
			},
		},
		{
			ID:       "c_0003",
			Brand:    "Toyota",
			Model:    "Celica",
			Year:     "1993",
			Colour:   "Red",
			Owner_Id: "0002",
			MalfunctionList: []MalfunctionAsset{
				{
					Desc:  "Low power",
					Price: 300.0,
				},
			},
		},
		{
			ID:       "c_0004",
			Brand:    "Mitsubishi",
			Model:    "Pajero",
			Year:     "1990",
			Colour:   "Mettalic Blue",
			Owner_Id: "0003",
			MalfunctionList: []MalfunctionAsset{
				{
					Desc:  "Broken suspension",
					Price: 420.0,
				},
			},
		},
		{
			ID:       "c_0005",
			Brand:    "Mitsubishi",
			Model:    "Pajero",
			Year:     "1990",
			Colour:   "Red",
			Owner_Id: "0003",
			MalfunctionList: []MalfunctionAsset{
				{
					Desc:  "Broken transmission",
					Price: 800.0,
				},
			},
		},
		{
			ID:       "c_0006",
			Brand:    "BMW",
			Model:    "E30",
			Year:     "1992",
			Colour:   "Gray",
			Owner_Id: "0003",
			MalfunctionList: []MalfunctionAsset{
				{
					Desc:  "Flat tire",
					Price: 230.0,
				},
				{
					Desc:  "Engine overheating",
					Price: 550.0,
				},
			},
		},
	}

	for _, personAsset := range personAssetList {
		personAssetJSON, err := json.Marshal(personAsset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(personAsset.ID, personAssetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	for _, carAsset := range carAssetList {
		carAssetJSON, err := json.Marshal(carAsset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(carAsset.ID, carAssetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadPersonAsset(ctx contractapi.TransactionContextInterface, personId string) (*PersonAsset, error) {
	personAssetJSON, err := ctx.GetStub().GetState(personId)
	if err != nil {
		return nil, fmt.Errorf("failed to read person from world state: %v", err)
	}
	if personAssetJSON == nil {
		return nil, fmt.Errorf("the person asset %s does not exist", personId)
	}

	var personAsset PersonAsset
	err = json.Unmarshal(personAssetJSON, &personAsset)
	if err != nil {
		return nil, err
	}

	return &personAsset, nil
}

func (s *SmartContract) ReadCarAsset(ctx contractapi.TransactionContextInterface, carId string) (*CarAsset, error) {
	carAssetJSON, err := ctx.GetStub().GetState(carId)
	if err != nil {
		return nil, fmt.Errorf("failed to read car from world state: %v", err)
	}
	if carAssetJSON == nil {
		return nil, fmt.Errorf("the person asset %s does not exist", carId)
	}

	var carAsset CarAsset
	err = json.Unmarshal(carAssetJSON, &carAsset)
	if err != nil {
		return nil, err
	}

	return &carAsset, nil
}

// AssetExists returns true when asset with given ID exists in world state
func (s *SmartContract) PersonAssetExists(ctx contractapi.TransactionContextInterface, personId string) (bool, error) {
	personAssetJSON, err := ctx.GetStub().GetState(personId)
	if err != nil {
		return false, fmt.Errorf("failed to read person from world state: %v", err)
	}

	return personAssetJSON != nil, nil
}

// UpdateAsset updates an existing asset in the world state with provided parameters.
func (s *SmartContract) TransferOwnership(ctx contractapi.TransactionContextInterface, carId string, newOwnerId string) error {
	carAsset, err := s.ReadCarAsset(ctx, carId)
	if err != nil {
		return err
	}

	exists, err := s.PersonAssetExists(ctx, newOwnerId)
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("the person %s does not exist", newOwnerId)
	}

	carAsset.Owner_Id = newOwnerId

	carAssetJSON, err := json.Marshal(carAsset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(carId, carAssetJSON)
}

// UpdateAsset updates an existing asset in the world state with provided parameters.
func (s *SmartContract) AddMulfunction(ctx contractapi.TransactionContextInterface, carId string, description string, price float64) error {
	carAsset, err := s.ReadCarAsset(ctx, carId)
	if err != nil {
		return err
	}

	malfunction := MalfunctionAsset{
		Desc:  description,
		Price: price,
	}

	carAsset.MalfunctionList = append(carAsset.MalfunctionList, malfunction)

	carAssetJSON, err := json.Marshal(carAsset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(carId, carAssetJSON)
}

// UpdateAsset updates an existing asset in the world state with provided parameters.
func (s *SmartContract) ChangeColour(ctx contractapi.TransactionContextInterface, carId string, colour string) error {
	carAsset, err := s.ReadCarAsset(ctx, carId)
	if err != nil {
		return err
	}

	carAsset.Colour = colour

	carAssetJSON, err := json.Marshal(carAsset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(carId, carAssetJSON)
}

// UpdateAsset updates an existing asset in the world state with provided parameters.
func (s *SmartContract) DoCarRepair(ctx contractapi.TransactionContextInterface, carId string) error {
	carAsset, err := s.ReadCarAsset(ctx, carId)
	if err != nil {
		return err
	}

	personAsset, err := s.ReadPersonAsset(ctx, carAsset.Owner_Id)
	if err != nil {
		return err
	}

	personCredit := personAsset.Credit
	totalRepairs := float64(0)

	for _, malfunction := range carAsset.MalfunctionList {
		totalRepairs += malfunction.Price
		if totalRepairs > personCredit {
			return fmt.Errorf("repairs can't be done because of insufficient funds")
		}
	}

	personAsset.Credit -= totalRepairs
	carAsset.MalfunctionList = []MalfunctionAsset{}

	carAssetJSON, err := json.Marshal(carAsset)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(carId, carAssetJSON)
	if err != nil {
		return err
	}

	personAssetJSON, err := json.Marshal(personAsset)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(carAsset.Owner_Id, personAssetJSON)
	if err != nil {
		return err
	}

	return nil
}

// // UpdateAsset updates an existing asset in the world state with provided parameters.
// func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, id string, color string, size int, owner string, appraisedValue int) error {
// 	exists, err := s.AssetExists(ctx, id)
// 	if err != nil {
// 		return err
// 	}
// 	if !exists {
// 		return fmt.Errorf("the asset %s does not exist", id)
// 	}

// 	// overwriting original asset with new asset
// 	asset := Asset{
// 		ID:             id,
// 		Color:          color,
// 		Size:           size,
// 		Owner:          owner,
// 		AppraisedValue: appraisedValue,
// 	}
// 	assetJSON, err := json.Marshal(asset)
// 	if err != nil {
// 		return err
// 	}

// 	return ctx.GetStub().PutState(id, assetJSON)
// }

// DeleteAsset deletes an given asset from the world state.
// func (s *SmartContract) DeleteAsset(ctx contractapi.TransactionContextInterface, id string) error {
// 	exists, err := s.AssetExists(ctx, id)
// 	if err != nil {
// 		return err
// 	}
// 	if !exists {
// 		return fmt.Errorf("the asset %s does not exist", id)
// 	}

// 	return ctx.GetStub().DelState(id)
// }

// GetAllAssets returns all assets found in world state
// func (s *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*Asset, error) {
// 	// range query with empty string for startKey and endKey does an
// 	// open-ended query of all assets in the chaincode namespace.
// 	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resultsIterator.Close()

// 	var assets []*Asset
// 	for resultsIterator.HasNext() {
// 		queryResponse, err := resultsIterator.Next()
// 		if err != nil {
// 			return nil, err
// 		}

// 		var asset Asset
// 		err = json.Unmarshal(queryResponse.Value, &asset)
// 		if err != nil {
// 			return nil, err
// 		}
// 		assets = append(assets, &asset)
// 	}

// 	return assets, nil
// }
