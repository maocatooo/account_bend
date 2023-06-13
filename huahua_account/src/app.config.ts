export default defineAppConfig({
  pages: [
    'pages/tabBar_book/index',
    'pages/tabBar_me/index',
    'pages/tabBar_add/index',
    'pages/me/book/index',
    'pages/me/tag/index',
  ],
  window: {
    backgroundTextStyle: 'light',
    navigationBarBackgroundColor: '#FFF',
    navigationBarTitleText: 'WeChat',
    navigationBarTextStyle: 'black'
  },
  tabBar: {
    color: "#ababab",
    selectedColor: "#000000",
    backgroundColor: "#FFF",
    borderStyle: "black",
    list: [
      {
        pagePath: "pages/tabBar_book/index",
        text: "book",
        iconPath: "static/img/tabbar/home.png",
        selectedIconPath: "static/img/tabbar/home.png"
      },
      {
        pagePath: "pages/tabBar_add/index",
        text: "",
        iconPath: "static/img/tabbar/add.png",
        selectedIconPath: "static/img/tabbar/add.png"
      },
      {
        pagePath: "pages/tabBar_me/index",
        text: "me",
        iconPath: "static/img/tabbar/me.png",
        selectedIconPath: "static/img/tabbar/me.png"
      }
    ],
  }
})
