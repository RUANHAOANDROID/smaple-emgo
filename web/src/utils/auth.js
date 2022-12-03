import Cookies from 'js-cookie'
const TokenKey = 'x-access-token'
var inFifteenMinutes = new Date(new Date().getTime() + 120 * 60 * 1000);
export function getToken() {
    return Cookies.get(TokenKey)
}

export function setToken(token) {
    return Cookies.set(TokenKey, token, { expires: inFifteenMinutes })
}

export function removeToken() {
    return Cookies.remove(TokenKey)
}

const mchIdKey = 'mchId'
export function getMchId() {
    return Cookies.get(mchIdKey)
}

export function setMchId(mchId) {
    return Cookies.set(mchIdKey, mchId, { expires: inFifteenMinutes })
}

export function removeMchId() {
    return Cookies.remove(mchIdKey)
}