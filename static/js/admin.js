
function edit(url, w, h) {
    if (!w){
        w = 600;
    }
    if (!h){
        h = 600;
    }
    $('<iframe id="editFrame" style="min-width:' + w + 'px; min-height:' + h + 'px"/>').attr("src", url).dialog({ autoOpen: true, modal: true, title: "编辑", width: w, height: h});
}

function del(url, callback) {
    if (confirm("是否确定执行操作?")) {
        $.get(url, function (msg) {
            if (msg == 'ok') {
                alert('操作成功！');
                refresh();
            } else if (msg == 'callback') {
                callback();
            } else {
                alert(msg);
            }
        });
    }
}

function req(url, callback) {
    if (confirm("是否确定执行操作?")) {
        
        if (url.indexOf("?")){
            url += "&_t=" + (new Date()).getTime(); 
        } else {
            url += "?_t=" + (new Date()).getTime();
        }
        
        $.getJSON(url, function (data) {
            //console.log(data);
            if (data['code'] == 0) {
                alert('操作成功！');
                refresh();
            } else if(data['code']>0){
                if (typeof callback == "function"){
                    callback(data["data"])
                }
                refresh();
            } else{
                alert(data["msg"]);
            }
        });
    }
}

function recover(url, callback) {
    $.get(url, function (msg) {
        if (msg == 'ok') {
            //alert('恢复成功！');
            refresh();
        } else if (msg == 'callback') {
            callback();
        } else {
            alert(msg);
        }
    });
}
function delAll(url, callback) {
    if (confirm("是否确定批量删除?")) {
        $.get(url, function (msg) {
            if (msg == 'ok') {
                //alert('操作成功！');
                refresh();
            } else if (msg == 'callback') {
                callback();
            } else {
                alert(msg);
            }
        });
    }
}
function recoverAll(url, callback) {
    if (confirm("是否确定批量恢复?")) {
        $.get(url, function (msg) {
            if (msg == 'ok') {
                //alert('操作成功！');
                refresh();
            } else if (msg == 'callback') {
                callback();
            } else {
                alert(msg);
            }
        });
    }
    return false;
}

function refresh() {
    location.href = location.href;
    location.reload(true);
}

if (typeof Modernizr == "undefined" || !Modernizr.input.placeholder) {
    $('[placeholder]').focus(function () {
        var input = $(this);
        if (input.val() == input.attr('placeholder')) {
            input.val('');
            input.removeClass('placeholder');
        }
    }).blur(function () {
            var input = $(this);
            if (input.val() == '' || input.val() == input.attr('placeholder')) {
                input.addClass('placeholder');
                input.val(input.attr('placeholder'));
            }
        }).blur();

    $('[placeholder]').parents('form').submit(function () {
        $(this).find('[placeholder]').each(function () {
            var input = $(this);
            if (input.val() == input.attr('placeholder')) {
                input.val('');
            }
        })
    });
}
