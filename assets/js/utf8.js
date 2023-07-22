function utf8Encode(inputStr) {
  var outputStr = "";
  
  for(var i = 0; i < inputStr.length; i++) {
    var temp = inputStr.charCodeAt(i);
    
    //0xxxxxxx
    if(temp < 128) {
      outputStr += String.fromCharCode(temp);
    }
    //110xxxxx 10xxxxxx
    else if(temp < 2048) {
      outputStr += String.fromCharCode((temp >> 6) | 192);
      outputStr += String.fromCharCode((temp & 63) | 128);
    }
    //1110xxxx 10xxxxxx 10xxxxxx
    else if(temp < 65536) {
      outputStr += String.fromCharCode((temp >> 12) | 224);
      outputStr += String.fromCharCode(((temp >> 6) & 63) | 128);
      outputStr += String.fromCharCode((temp & 63) | 128);
    }
    //11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
    else {
      outputStr += String.fromCharCode((temp >> 18) | 240);
      outputStr += String.fromCharCode(((temp >> 12) & 63) | 128);
      outputStr += String.fromCharCode(((temp >> 6) & 63) | 128);
      outputStr += String.fromCharCode((temp & 63) | 128);
    }
  }
  
  return outputStr;
}

function utf8Decode(inputStr) {
  var outputStr = "";
  var code1, code2, code3, code4;
  
  for(var i = 0; i < inputStr.length; i++) {
    code1 = inputStr.charCodeAt(i);
    
    if(code1 < 128) {
      outputStr += String.fromCharCode(code1);
    }
    else if(code1 < 224) {
      code2 = inputStr.charCodeAt(++i);
      outputStr += String.fromCharCode(((code1 & 31) << 6) | (code2 & 63));
    }
    else if(code1 < 240) {
      code2 = inputStr.charCodeAt(++i);
      code3 = inputStr.charCodeAt(++i);
      outputStr += String.fromCharCode(((code1 & 15) << 12) | ((code2 & 63) << 6) | (code3 & 63));
    }
    else {
      code2 = inputStr.charCodeAt(++i);
      code3 = inputStr.charCodeAt(++i);
      code4 = inputStr.charCodeAt(++i);
      outputStr += String.fromCharCode(((code1 & 7) << 18) | ((code2 & 63) << 12) |((code3 & 63) << 6) | (code2 & 63));
    }
  }
  
  return outputStr;
}