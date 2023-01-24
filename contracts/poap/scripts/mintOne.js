// scripts/mint.js
const { ethers } = require("hardhat");
const { CONTRACT_ADDRESS } = require("./config");

async function main() {
  console.log("Getting the FlaqPOAP contract...\n");
  const contractAddress = CONTRACT_ADDRESS;
  const flaqPoap = await ethers.getContractAt("FlaqPOAP", contractAddress);
  const signers = await ethers.getSigners();

  const users = [];

  console.log("Running");
  let tx = await flaqPoap.mintCollectionNFT(
    "0xedAAE8847970b79f07AD01332ce1142089eB2E7F",
    "https://flaq-assets.s3.ap-south-1.amazonaws.com/poap1/jsons/shreya.json"
  );
  console.log(tx.hash);
  await tx.wait(); // wait for this tx to finish to avoid nonce issues
  console.log("---");
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exitCode = 1;
  });
