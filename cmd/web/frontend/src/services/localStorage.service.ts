import { types } from "../../wailsjs/go/models";

export const persistKey = (keys: types.Keys): types.Keys => {
  localStorage.setItem("keys", JSON.stringify(keys));
  return keys;
};

export const persistKekType = (kekType: string): string => {
  localStorage.setItem("kekType", kekType);
  return kekType;
};

export const readKey = (): types.Keys => {
  return localStorage.getItem("keys")
    ? JSON.parse(localStorage.getItem("keys") as string)
    : null;
};

export const readKekType = (): string => {
  return localStorage.getItem("kekType") || "";
};

export const deleteKey = (): void => localStorage.removeItem("keys");
export const deleteKekType = (): void => localStorage.removeItem("kekType");
