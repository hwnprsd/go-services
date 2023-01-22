// scripts/mint.js
const { ethers } = require("hardhat");

async function main() {
  console.log("Getting the non fun token contract...\n");
  const contractAddress = "0x5FbDB2315678afecb367f032d93F642f64180aa3";
  const nonFunToken = await ethers.getContractAt(
    "NonFunToken",
    contractAddress
  );
  const signers = await ethers.getSigners();
  const contractOwner = signers[0].address;
  const newOwner = "0xa0Ee7A142d267C1f36714E4a8F75612F20a79720";

  let tx = await nonFunToken.tokenURI("11");
  console.log(tx);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exitCode = 1;
  });
