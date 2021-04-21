const SCALE = 10;

const pointsColor = '#000000';
const colors = [
    '#F7E017',
    '#ED6E00',
    '#E60094',
    '#7800b9',
    '#2905A1',
    '#008CCC',
    '#00B394',
    '#00AB59',
    '#F54029',
    '#FC8CA1',
    '#E66BC2',
    '#B878BF',
    '#73B5E0',
];
let lastColor = 0;

function getClusterColor() {
    let color = colors[lastColor];
    lastColor++;
    lastColor %= colors.length;
    return color;
}

let areaId = -1;
let points = [];
let clusters = [];

function clearAll() {
    ClearCanvas(canvas);
    points = [];
    clusters = [];
    lastColor = 0;
}

function printPoint(ctx, x, y, color) {
    ctx.fillStyle = color;
    ctx.fillRect(x, y, 1, 1);
}

function addPoint(ctx, x, y, color) {
    printPoint(ctx, x, y, color);
    points.push({x: x, y: y});
}

function printCluster(ctx, x, y, color) {
    ctx.fillStyle = color;
    ctx.fillRect(x - 1, y, 1, 1);
    ctx.fillRect(x, y - 1, 1, 1);
    ctx.fillRect(x + 1, y, 1, 1);
    ctx.fillRect(x, y + 1, 1, 1);
    ctx.fillRect(x, y, 1, 1);
}

function addCluster(ctx, x, y, color) {
    printCluster(ctx, x, y, color);
    clusters.push({x: x, y: y});
}

function train(distId, byStep, maxAge) {
    ApiAddPoints(areaId, points, () => {
        ApiAddClusters(areaId, clusters, () => {
            if (byStep && points.length > 0) {
                ApiGetArea(areaId, distId, (data) => {
                    clearAll();

                    data.clusters.forEach((cluster) => {
                        let color = getClusterColor();
                        printCluster(ctx, cluster.x, cluster.y, color);
                        cluster.points.forEach((point) => {
                            printPoint(ctx, point.x, point.y, color);
                        });
                    });
                });
            } else {
                ApiTrain(areaId, distId, byStep, maxAge, (data) => {
                    clearAll();

                    data.clusters.forEach((cluster) => {
                        let color = getClusterColor();
                        printCluster(ctx, cluster.x, cluster.y, color);
                        cluster.points.forEach((point) => {
                            printPoint(ctx, point.x, point.y, color);
                        });
                    });
                })
            }
        });
    });
}

function clearArea() {
    ApiClearArea(areaId);
    clearAll();
}

function leftClick(event) {
    event.preventDefault();
    addPoint(ctx, Math.trunc(event.offsetX / SCALE), Math.trunc(event.offsetY / SCALE), pointsColor);
}

function rightClick(event) {
    event.preventDefault();
    addCluster(ctx, Math.trunc(event.offsetX / SCALE), Math.trunc(event.offsetY / SCALE), getClusterColor());
}

// Main code
let canvas = CanvasInit('canvas', SCALE);
let ctx = canvas.getContext('2d');
canvas.addEventListener('click', leftClick);
canvas.addEventListener('contextmenu', rightClick);

ApiAddArea((data) => {
    areaId = data.id;
    console.log("Area ID:", areaId);
});

function randomInt(max) {
    return Math.floor(Math.random() * max);
}

// Test data
function addTestData() {
    for (let i = 0; i < 50; i++) {
        addPoint(ctx, randomInt(canvas.width / SCALE), randomInt(canvas.height / SCALE), pointsColor);
    }
    for (let i = 0; i < 5; i++) {
        addCluster(ctx, randomInt(canvas.width / SCALE), randomInt(canvas.height / SCALE), getClusterColor());
    }
}

