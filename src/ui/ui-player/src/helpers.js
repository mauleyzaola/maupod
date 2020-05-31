const msToString = ms => (new Date(ms)).toGMTString().split(' ')[4];

const secondsToDate = seconds => {
    const epoch = new Date(1970,0,1);
    epoch.setSeconds(seconds);
    return epoch;
}

export {
    msToString,
    secondsToDate,
}
