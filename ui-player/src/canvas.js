
let   CANVAS_WIDTH = 1920;
const CANVAS_HEIGHT = 150;

// these are the colors when track has been played
const darkRed = parseInt('2e', 16);
const darkGreen = parseInt('84', 16);
const darkBlue = parseInt('9e', 16);

// these are the colors from the original spectrum
const black = 0;
const red = 70;
const green = 194;
const blue = 230;

const paintedColors = { red, green, blue };
const blurredColors = { red: darkRed, green: darkGreen, blue: darkBlue };

let ctx;

// loadDataItems will draw a canvas and instantiate the context for it from an image source
const loadCanvasImage = ({canvas, src, width}) => {
    CANVAS_WIDTH = width;
    try{
        if(!canvas) throw Error('missing canvas parameter');
        if(!src) throw Error('missing src parameter');
        const img = new Image()
        img.crossOrigin = '';
        img.src = src;
        canvas.width = CANVAS_WIDTH - 10;
        ctx = canvas.getContext('2d');
        
        img.addEventListener('load', () => {
            ctx.drawImage(img, 0, 0, CANVAS_WIDTH, CANVAS_HEIGHT);
        })
    }catch (e){
        console.error(e);
    }
}

// blurRange will blur a section of the canvas in from x1 to x2 across the canvas width
const blurRange = ({dataItems, x1, x2}) => {
    for (let y = 0; y < CANVAS_HEIGHT; y++) {
        for (let x = x1; x <= CANVAS_WIDTH; x++) {
            const values = getColorIndicesForCoord(x, y, CANVAS_WIDTH);
            setColorIndices(dataItems, values, x < x2 ? blurredColors : paintedColors );
        }
    }
    return dataItems;
}

// applyBlur will simulate a blur effect for an audio canvas
const applyBlur = ({x1,x2}) => {
    if(!ctx) return;
    let imageData = ctx.getImageData(0,0,CANVAS_WIDTH,CANVAS_HEIGHT);
    let dataItems = imageData.data;
    blurRange({dataItems, x1, x2});
    imageData.data.set(dataItems);
    ctx.putImageData(imageData, 0, 0);
}

const getColorIndicesForCoord = (x, y, width) => {
    const red = y * (width * 4) + x * 4;
    return [red, red + 1, red + 2, red + 3];
};

const isBlack = ([r,g,b,a]) => r+g+b+a === black;

const setColorIndices = (items, [r,g,b,a], {red, green, blue}) => {
    let values = [items[r],items[g],items[b]];
    if(isBlack(values)){
        return values;
    }
    // if the color is not black, we assume we need to blur
    items[r] = red;
    items[g] = green;
    items[b] = blue;
    // alpha remains the same, so image is still transparent and can play well with any background color
    return values
}

export {
    applyBlur,
    CANVAS_HEIGHT,
    CANVAS_WIDTH,
    loadCanvasImage,
}