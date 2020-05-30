const secondsToDate = seconds => {
    const epoch = new Date(1970,0,1);
    epoch.setSeconds(seconds);
    return epoch;
}

export {
    secondsToDate,
}
