// groupOnFirstChar accepts an array of strings
// and returns an object of arrays taking the first letter of each element
// items are assumed to be already sorted
const groupOnFirstChar = items => {
    let rows = {};
    for(let i=0; i < items.length; i++){
        const value = items[i];
        if(typeof value === 'undefined') continue;
        if(value.length === 0) continue;
        const key = value[0];
        let item = rows[key];
        if(!item){
            item = [];
            rows[key] = item;
        }
        item.push(value);
    }
    let result = [];
    for(let k in rows){
        result.push({
            letter: k,
            items: rows[k],
        })
    }
    return result;
}

const msToString = ms => (new Date(ms)).toGMTString().split(' ')[4];

const secondsToDate = seconds => {
    const epoch = new Date(1970,0,1);
    epoch.setSeconds(seconds);
    return epoch;
}
const msToDate = ms => {
    const epoch = new Date(1970,0,1);
    epoch.setMilliseconds(ms);
    return epoch;
}

export {
    groupOnFirstChar,
    msToDate,
    msToString,
    secondsToDate,
}
