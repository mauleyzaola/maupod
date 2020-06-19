// groupOnFirstChar accepts an array of strings
// and returns an object of arrays taking the first letter of each element
// items are assumed to be already sorted
const groupOnFirstChar = items => {
    let result = {};
    for(let i=0; i < items.length; i++){
        const value = items[i];
        console.log(`value: ${value}`)
        if(value.length === 0) continue;
        const key = value[0];
        let item = result[key];
        if(!item){
            item = [];
            result[key] = item;
        }
        item.push(value);
    }
    return result;
}

const msToString = ms => (new Date(ms)).toGMTString().split(' ')[4];

const secondsToDate = seconds => {
    const epoch = new Date(1970,0,1);
    epoch.setSeconds(seconds);
    return epoch;
}

export {
    groupOnFirstChar,
    msToString,
    secondsToDate,
}
