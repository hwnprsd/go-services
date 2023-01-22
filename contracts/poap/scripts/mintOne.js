// scripts/mint.js
const { ethers } = require("hardhat");
const { CONTRACT_ADDRESS } = require("./config");

async function main() {
  console.log("Getting the non fun token contract...\n");
  const contractAddress = CONTRACT_ADDRESS;
  const flaqPoap = await ethers.getContractAt("FlaqPOAP", contractAddress);
  const signers = await ethers.getSigners();
  const contractOwner = signers[0].address;
  const newOwner = "0x0Dac92D471EC31e8F4004ae87191344D974F108d";

  let tx = await flaqPoap.mintCollectionNFT(
    "0xC6BaBae03FcfEe9fB709EAC8500eCbAfbDDC461d",
    "https://flaq-assets.s3.ap-south-1.amazonaws.com/poap1/jsons/Ankit.json"
  );
  await tx.wait(); // wait for this tx to finish to avoid nonce issues
  console.log(tx);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exitCode = 1;
  });
