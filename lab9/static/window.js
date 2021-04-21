function CreateSaveLink(canvas) {
    let bar = document.getElementById('btn-bar');
    if (bar) {
        let link = document.createElement('a');
        link.innerHTML = 'Save ' + canvas.id;
        //link.href = '#';
        //link.className = 'save';
        link.addEventListener('click', function () {
            link.href = canvas.toDataURL();
            link.download = 'canvas.png';
        }, false);
        bar.appendChild(link);
    }
}

function CropCanvas(canvas) {
    let ctx = canvas.getContext('2d');
    let imageData = ctx.getImageData(0, 0, canvas.width, canvas.height);
    canvas.width = canvas.offsetWidth;
    canvas.height = canvas.offsetHeight;
    ctx.putImageData(imageData, 0, 0);

    ctx.scale(canvas.scale, canvas.scale);
}

function ResizeCanvas(canvas) {
    let ctx = canvas.getContext('2d');

    canvas.style.width = '';
    canvas.style.height = '';

    let canvas_ratio = canvas.width / canvas.height;
    let offset_ratio = canvas.offsetWidth / canvas.offsetHeight;
    if (canvas_ratio > offset_ratio) {
        canvas.style.height = canvas.offsetWidth / canvas_ratio + 'px';
    } else {
        canvas.style.width = canvas.offsetHeight * canvas_ratio + 'px';
    }

    ctx.scale(canvas.scale, canvas.scale);
}

function Resize() {
    Resize.canvases.forEach(function (canvas) {
        CropCanvas(canvas);
    });
    Resize.scalable_canvases.forEach(function (canvas) {
        ResizeCanvas(canvas);
    });
    if (Resize.callback) {
        Resize.callback();
    }
}

Resize.canvases = [];
Resize.scalable_canvases = [];
Resize.callback = null;
window.onresize = Resize;

function CanvasInit(id, scale, width, height) {
    let canvas = document.getElementById(id);
    CreateSaveLink(canvas);
    if (arguments.length <= 2) {
        canvas.width = canvas.offsetWidth;
        canvas.height = canvas.offsetHeight;
        canvas.scale = scale;
        CropCanvas(canvas);
        Resize.canvases.push(canvas);
    } else {
        canvas.width = width;
        canvas.height = height;
        canvas.scale = scale;
        ResizeCanvas(canvas);
        Resize.scalable_canvases.push(canvas);
    }
    ClearCanvas(canvas);
    return canvas;
}

function ClearCanvas(canvas) {
    let ctx = canvas.getContext('2d');
    ctx.clearRect(0, 0, canvas.width, canvas.height)
    ctx.fillStyle = '#eee';
    for (let x = 0; x < canvas.width; x++) {
        for (let y = x % 2; y < canvas.height; y += 2) {
            ctx.fillRect(x, y, 1, 1)
        }
    }
}

function LoadImage(old_canvas, src, new_canvas, callback) {
    let img = new Image();
    let args = arguments.length;
    img.onload = function () {
        old_canvas.getContext('2d').drawImage(img, 0, 0);
        if (args === 4) {
            callback(old_canvas, new_canvas);
        }
    };
    img.src = src;
}
