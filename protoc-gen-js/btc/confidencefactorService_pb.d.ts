export class TXConfidence {
  constructor ();
  getAgeMillis(): number;
  setAgeMillis(a: number): void;
  getReceiveCount(): number;
  setReceiveCount(a: number): void;
  getConfidence(): number;
  setConfidence(a: number): void;
  getTxhash(): string;
  setTxhash(a: string): void;
  getTxurl(): string;
  setTxurl(a: string): void;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => TXConfidence;
}

export class TransactionConfidenceEndpointRequest {
  constructor ();
  getNetwork(): {};
  setNetwork(a: {}): void;
  getTxhash(): string;
  setTxhash(a: string): void;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => TransactionConfidenceEndpointRequest;
}
