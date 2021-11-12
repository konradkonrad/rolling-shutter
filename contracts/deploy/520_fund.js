const { ethers } = require("hardhat");

module.exports = async function (hre) {
  const fundValue = hre.deployConf.fundValue;
  if (fundValue == "") {
    console.log("fund: not doing any funding");
    return;
  }

  const [owner] = await ethers.getSigners();
  const decryptor = await hre.getDecryptorAddresses();
  const keypers = await hre.getKeyperAddresses();
  const addresses = decryptor.concat(keypers);
  const value = ethers.utils.parseEther(fundValue);
  console.log(
    "fund: funding %s adresses with %s eth",
    addresses.length,
    fundValue
  );

  const txs = [];
  for (const a of addresses) {
    const tx = await owner.sendTransaction({
      to: a,
      value: value,
    });
    txs.push(tx);
  }
  for (const tx of txs) {
    await tx.wait();
  }
};
