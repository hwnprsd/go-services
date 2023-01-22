// scripts/mint.js
const { ethers } = require("hardhat");
const { CONTRACT_ADDRESS } = require("./config");

async function main() {
  console.log("Getting the non fun token contract...\n");
  const contractAddress = CONTRACT_ADDRESS;
  const nonFunToken = await ethers.getContractAt("FlaqPOAP", contractAddress);
  const signers = await ethers.getSigners();
  const contractOwner = signers[0].address;
  const newOwner = "0xa0Ee7A142d267C1f36714E4a8F75612F20a79720";

  // let tx = await nonFunToken.mintCollectionNFT(newOwner, "https://google.com");
  console.log(await nonFunToken.tokenURI("0"));
  // await tx.wait(); // wait for this tx to finish to avoid nonce issues
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exitCode = 1;
  });
