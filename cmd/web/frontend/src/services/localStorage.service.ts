import { types } from "../../wailsjs/go/models";

export const persistKey = (keys: types.Keys): void => {
  localStorage.setItem("keys", JSON.stringify(keys));
};

export const readKey = (): types.Keys => {
  return localStorage.getItem("keys")
    ? JSON.parse(localStorage.getItem("keys") as string)
    : null;
};

export const deleteKey = (): void => localStorage.removeItem("keys");
