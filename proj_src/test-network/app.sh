#!/usr/bin/env bash

export PATH=${PWD}/../bin:$PATH

export FABRIC_CFG_PATH=$PWD/../config/

export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org3MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer1.org3.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org3.example.com/users/User1@org3.example.com/msp
export CORE_PEER_ADDRESS=localhost:11061

optii() {
    echo "Enter car id:  "
    read inp

    query='{"Args":["ReadCarAsset",'\"$inp\"']}'

    peer chaincode query -C mychannel -n basic -c $query
}

optii2() {
    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"InitLedger","Args":[]}'
}

optii3() {
    echo "Enter car id:  "
    read inp_car
    echo "Enter new owner id:  "
    read inp_owner
    echo "Price:  "
    read inp_price
    echo "Pay for repairs: [Press enter for yes]"
    read inp_rep

    if [$inp_rep == ""]
    then 
        inp_rep_bool="true"
    else
        inp_rep_bool="false"
    fi

    echo $inp_rep_bool

    query='{"function":"TransferOwnership","Args":['\"$inp_car\"','\"$inp_owner\"','\"$inp_price\"','\"$inp_rep_bool\"']}'

    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:11051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt" -c $query
}

optii4() {
    echo "Enter person id:  "
    read inp

    query='{"Args":["ReadPersonAsset",'\"$inp\"']}'

    peer chaincode query -C mychannel -n basic -c $query
}

optii5() {
    echo "Enter car id:  "
    read inp_car
    echo "Enter new colour:  "
    read inp_col

    query='{"function":"ChangeColour","Args":['\"$inp_car\"','\"$inp_col\"']}'

    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:11051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt" -c $query
}

optii6() {
    echo "Enter car id:  "
    read inp_car
    echo "Enter malfunction description:  "
    read inp_desc
    echo "Enter malfunction price:  "
    read inp_price

    query='{"function":"AddMulfunction","Args":['\"$inp_car\"','\"$inp_desc\"','\"$inp_price\"']}'

    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:11051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt" -c $query
}

optii7() {
    echo "Enter car id:  "
    read inp_car

    query='{"function":"DoCarRepair","Args":['\"$inp_car\"']}'

    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:11051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt" -c $query
}

mainmenu() {
    echo -ne "
MAIN MENU
1) Car Asset
2) Initialize ledger
3) Transfer Ownership
4) Person Asset
5) Change Colour
6) Add Malfunction
7) Repair Car
0) Exit
Choose an option:  "
    read -r ans
    case $ans in
    7)
        optii7
        mainmenu
        ;;
    6)
        optii6
        mainmenu
        ;;
    5)
        optii5
        mainmenu
        ;;
    4)
        optii4
        mainmenu
        ;;
    3)
        optii3
        mainmenu
        ;;
    2)
        optii2
        mainmenu
        ;;
    1)
        optii
        mainmenu
        ;;
    0)
        echo "Bye bye."
        exit 0
        ;;
    *)
        echo "Wrong option."
        exit 1
        ;;
    esac
}

# ./network.sh up createChannel -c mychannel -ca
# ./network.sh deployCC -ccn basic -ccp ../asset-transfer-cars/ -ccl go


mainmenu