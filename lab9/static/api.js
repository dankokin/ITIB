function Ajax(type, url, data, callback) {
    type = type.toUpperCase();
    $.ajax({
        type: type,
        url: url,
        data: type === 'GET' ? data : JSON.stringify(data),
        dataType: type === 'GET' ? '' : 'json',
    })
        .done((answer) => {
            if (100 <= answer.status && answer.status < 200) {
                if (callback) {
                    callback(answer.data);
                }
            } else {
                console.error(answer.message);
            }
        })
        .fail((error) => {
            console.error(error);
        })
}

function Get(url, data, callback) {
    Ajax('GET', url, data, callback)
}

function Post(url, data, callback) {
    Ajax('POST', url, data, callback)
}

function Delete(url, data, callback) {
    Ajax('DELETE', url, data, callback)
}

function ApiAddArea(callback) {
    Post("/api/area", {}, callback)
}

function ApiAddPoints(areaId, points, callback) {
    Post("/api/point", {id: areaId, points: points}, callback)
}

function ApiAddClusters(areaId, clusters, callback) {
    Post("/api/cluster", {id: areaId, clusters: clusters}, callback)
}

function ApiTrain(areaId, distanceId, byStep, maxAge, callback) {
    Post("/api/train", {id: areaId, dist_id: distanceId, by_step: byStep, max_age: maxAge}, callback)
}

function ApiGetArea(areaId, distanceId, callback) {
    Get("/api/area", {id: areaId, dist_id: distanceId}, callback)
}

function ApiClearArea(areaId, callback) {
    Delete("/api/area", {id: areaId}, callback)
}
