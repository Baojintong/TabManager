chrome.runtime.onInstalled.addListener(function() {
    chrome.contextMenus.create({
        id: "saveSession",
        title: "Save Session",
        contexts: ["all"]
    });
});

chrome.contextMenus.onClicked.addListener(function(info, tab) {
    if (info.menuItemId === "saveSession") {
        chrome.tabs.query({}, function(tabs) {
            const sessionData = tabs.map(function(tab) {
                return {
                    title: tab.title,
                    url: tab.url,
                    iconUrl:tab.favIconUrl
                };
            });
            // 发送POST请求
            fetch("http://localhost:12315/tabs", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(sessionData)
            })
                .then(function(response) {
                    if (response.ok) {
                        chrome.tabs.query({}, function(tabs) {
                            const tabIds = tabs.map(function(tab) {
                                return tab.id;
                            });
                            chrome.tabs.remove(tabIds);
                        });
                    } else {
                        console.log("Error:",response.status);
                    }
                })
                .catch(function(error) {
                    console.log("Error:", error);
                });
        });
    }
});