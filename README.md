# cosmos-sdk-demo
A repo for an automated process to initialize and make transactions on COSMOS-SDK.
1. Clone and install https://github.com/cosmos/cosmos-sdk
2. Clone ad install https://github.com/pigfox/cosmos-sdk-demo
3. You want this folder structure:
```
$ tree -L 1
.
├── cosmos-sdk
├── cosmos-sdk-demo
```
4. To run this application simply type `./cosmos-sdk-demo$ go run .`

Debug!
To switch between the initial genesis file structure and an experimental go to 
addGenesisFile.go.
Comment/Uncomment linnes 31/32
```
genesisJson := getGenesisJSONX(gp)
//genesisJson := getGenesisJSON0(gp)
```
