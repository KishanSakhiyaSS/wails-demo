import { GetUser } from "../wailsjs/wailsjs/go/handlers/UserHandler";

export async function getUser() {
  return await GetUser();
}
