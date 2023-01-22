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
  let tx = await nonFunToken.balanceOf(signers[0].address);
  console.log(tx);
  // await tx.wait(); // wait for this tx to finish to avoid nonce issues
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exitCode = 1;
  });
