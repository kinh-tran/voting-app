# voting-app

## Get the Requirements
- Linux/MacOS
  - https://hyperledger-fabric.readthedocs.io/en/latest/prereqs.html

## Clone the Repo
```
git clone https://github.com/kinh-tran/voting-app.git
cd voting-app
```

## Set Environment Variables
```
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051
```

## Deploy the Network
```
cd test-network
./network.sh up createChannel -c mychannel -ca
./network.sh deployCC -ccn vote -ccp ../chaincode -ccl go
```

## Run the App
```
cd ..
go run ./cmd/
```

## On your browser
- Navigate to http://localhost:4445 
