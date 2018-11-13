import * as grpcWeb from 'grpc-web';
import {
  Address,
  AddressBalanceEndpointRequest,
  AddressEndpointRequest,
  AddressFullEndpointRequest,
  AddressKeychain,
  GenerateAddressEndpointRequest,
  GenerateMultisigAddressEndpointRequest} from './addressService_pb';

export class AddressServiceClient {
  constructor (hostname: string,
               credentials: {},
               options: { [s: string]: {}; });

  addressBalanceEndpoint(
    request: AddressBalanceEndpointRequest,
    metadata: grpcWeb.Metadata,
    callback: (err: grpcWeb.Error,
               response: Address) => void
  ): grpcWeb.ClientReadableStream;

  addressEndpoint(
    request: AddressEndpointRequest,
    metadata: grpcWeb.Metadata,
    callback: (err: grpcWeb.Error,
               response: Address) => void
  ): grpcWeb.ClientReadableStream;

  addressFullEndpoint(
    request: AddressFullEndpointRequest,
    metadata: grpcWeb.Metadata,
    callback: (err: grpcWeb.Error,
               response: Address) => void
  ): grpcWeb.ClientReadableStream;

  generateAddressEndpoint(
    request: GenerateAddressEndpointRequest,
    metadata: grpcWeb.Metadata,
    callback: (err: grpcWeb.Error,
               response: AddressKeychain) => void
  ): grpcWeb.ClientReadableStream;

  generateMultisigAddressEndpoint(
    request: GenerateMultisigAddressEndpointRequest,
    metadata: grpcWeb.Metadata,
    callback: (err: grpcWeb.Error,
               response: AddressKeychain) => void
  ): grpcWeb.ClientReadableStream;

}
