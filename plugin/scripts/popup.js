let tabs = await chrome.tabs.query({});


let collator = new Intl.Collator();
tabs.sort((a, b) => collator.compare(a.title, b.title));
let dataList = [];
for (let tab of tabs) {
    let title = tab.title.split('-')[0].trim();
    let url = tab.url;
    let iconUrl = tab.favIconUrl;

    let data = {};
    data['title'] = title;
    data['url'] = url;
    data['iconUrl'] = iconUrl;
    dataList.push(data);
}

let button = document.querySelector('button');
button.addEventListener('click', async () => {
    const jsonData = JSON.stringify(dataList);
    $.post("http://localhost:12315/tabs",jsonData,
        function (data, status) {
            alert("Data: " + data + "\nStatus: " + status);
        }, "json");
    alert("保存完成");
});