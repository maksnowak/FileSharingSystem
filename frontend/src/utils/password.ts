export const hashPassword = async (password: string, salt: string) => {
  let saltedPassword = password + salt;
  let hashBuffer = await crypto.subtle.digest('SHA-256', new TextEncoder().encode(saltedPassword));
  return Array.prototype.map.call(new Uint8Array(hashBuffer), x=>(('00'+x.toString(16)).slice(-2))).join('');
}
