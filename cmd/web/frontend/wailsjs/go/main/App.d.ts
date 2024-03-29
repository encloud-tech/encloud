// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {types} from '../models';

export function FetchConfig():Promise<types.ConfigResponse>;

export function GenerateKeyPair(arg1:string):Promise<types.GenerateKeyPairResponse>;

export function List(arg1:string):Promise<types.ListContentResponse>;

export function ListKeys():Promise<types.ListKeysResponse>;

export function RestoreDefaultConfig():Promise<types.ConfigResponse>;

export function RetrieveByUUID(arg1:string,arg2:string,arg3:string,arg4:string):Promise<types.RetrieveByCIDContentResponse>;

export function RetrieveSharedContent(arg1:string,arg2:string,arg3:string,arg4:string,arg5:string):Promise<types.SharedResponse>;

export function SelectDirectory():Promise<string>;

export function SelectFile():Promise<string>;

export function Share(arg1:string,arg2:string,arg3:string,arg4:string):Promise<types.RetrieveByCIDContentResponse>;

export function StoreConfig(arg1:types.ConfYaml):Promise<types.ConfigResponse>;

export function Upload(arg1:string,arg2:string,arg3:string,arg4:string):Promise<types.UploadContentResponse>;
