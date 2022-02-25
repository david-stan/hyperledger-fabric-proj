## Usage

- Since interaction with the chaincode is through peer api, initialization and console usage is done by 
/proj_src/test-network/app.sh

- Just uncomment two commands at the bottom to setup the network and install chaincode, at least the first time.

- User1 of the organization 3 is already being set up in the beginning of the file since it is a requirement of the project. Peer 1 of the organization 3 is used in this instance.

- "Car Asset" and "Person Asset" are used to review changes after invoke commands.


- Known bugs: 
  - While entering description for adding car malfunction, no spaces must be entered, since bash has problems with spaces and crashes the menu. TBD
