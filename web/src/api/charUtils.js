import vPinyin from './vue-py.js';

export function getCamelChars (ch) {
    let pinYin = vPinyin.chineseToPinYin(ch);
    let SX = '';
    for (var i = 0; i < pinYin.length; i++) {
        var c = pinYin.charAt(i);
        if (/^[A-Z]+$/.test(c)) {
            SX += c;
        }
    }
    return SX;
}