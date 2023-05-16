import { Slip10RawIndex } from "@cosmjs/crypto";

export function makeCosmoshubPath(slip44 ,a) {
  return [
    Slip10RawIndex.hardened(44),
    Slip10RawIndex.hardened(slip44),
    Slip10RawIndex.hardened(0),
    Slip10RawIndex.normal(0),
    Slip10RawIndex.normal(a),
  ];
}
