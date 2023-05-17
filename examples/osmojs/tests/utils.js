import BigNumber from 'bignumber.js';

export const calcShareOutAmount = (
    poolInfo,
    coinsNeeded
) => {
  return poolInfo.poolAssets
    .map(({ token }, i) => {
      const tokenInAmount = new BigNumber(coinsNeeded[i].amount);
      const totalShare = new BigNumber(poolInfo.totalShares.amount);
      const totalShareExp = totalShare.shiftedBy(-18);
      const poolAssetAmount = new BigNumber(token.amount);
      
      return tokenInAmount
        .multipliedBy(totalShareExp)
        .dividedBy(poolAssetAmount)
        .shiftedBy(18)
        .decimalPlaces(0, BigNumber.ROUND_HALF_UP)
        .toString();
    })
    .sort((a, b) => (new BigNumber(a).lt(b) ? -1 : 1))[0];
};

export const calcAmountWithSlippage = (
  amount,
  slippage
) => {
  const remainingPercentage = new BigNumber(100).minus(slippage).div(100);
  return new BigNumber(amount).multipliedBy(remainingPercentage).toString();
};

export const daysToSeconds = (days) => {
  return (Number(days) * 24 * 60 * 60).toString();
};

export const waitUntil = (date, timeout = 90000) => {
  return new Promise((resolve) => {
    const delay = date.getTime() - Date.now();
    if (delay > timeout) {
      throw new Error('Timeout to wait until date');
    }
    setTimeout(resolve, delay + 1000);
  });
};
