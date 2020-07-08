$(function () {
    if (IsPC()) {
        $(".direction_button").on({
            mousedown: function () {
                var id = $(this).attr("id");
                press(id);
            },
            mouseup: function () {
                var id = $(this).attr("id");
                release(id);
            },
            mouseleave: function () {
                var id = $(this).attr("id");
                release(id);
            }
        })
    }
    else {
        $(".direction_button").on({
            touchstart: function () {
                var id = $(this).attr("id");
                press(id);
            },
            touchend: function () {
                var id = $(this).attr("id");
                release(id);
            }
        })
    }
});


function IsPC() {
    var userAgentInfo = navigator.userAgent;
    var Agents = ["Android", "iPhone", "SymbianOS", "Windows Phone", "iPod"];
    var flag = true;
    for (var v = 0; v < Agents.length; v++) {
        if (userAgentInfo.indexOf(Agents[v]) > 0) {
            flag = false;
            break;
        }
    }
    if (window.screen.width >= 768) {
        flag = true;
    }
    return flag;
}

var vm = new Vue({
    el: "#http_alert",
    data: {
        code: "",
        action: ""
    }
})

function press(id) {
    $.post("/press",
        {
            code: id
        },
        function (data, status) {
            vm.$data.code = data.code
            vm.$data.action = data.action
        });
}

function release(id) {
    $.post("/release",
        {
            code: id
        },
        function (data, status) {
            vm.$data.code = data.code
            vm.$data.action = data.action
        });
}