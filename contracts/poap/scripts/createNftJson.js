const fs = require("fs");

const createObject = (name) => ({
  description: `Proof that ${name} attended the FLAQ x Rotaract Bangalore Junction session on 'Dive into Web3'`,
  external_url: "https://flaq.club",
  image: `https://flaq-assets.s3.ap-south-1.amazonaws.com/poap1/images/${name}.gif`,
  name: "Flaq x Rotaract Bangalore Junction",
  attributes: [
    {
      trait_type: "Level",
      value: "Dive into web3",
    },
  ],
});

async function main() {
  const names = ["Ashwin", "Muaaz", "Ankit", "Sarthak"];
  for (const name of names) {
    const json = createObject(name);
    fs.writeFileSync(`nfts/${name}.json`, JSON.stringify(json, null, 5));
  }
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exitCode = 1;
  });
