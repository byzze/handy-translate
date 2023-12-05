// @ts-check
// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT


window.go = window.go || {};
window.go.main = {
    App: {

		/**
		 * App.CaptureSelectedScreen
		 * CaptureSelectedScreen 截取选中的区域
		 * @param startX {number}
 * @param startY {number}
 * @param width {number}
 * @param height {number}
		 * @returns {Promise<void>}
		 **/
	    CaptureSelectedScreen: function(startX, startY, width, height) { return wails.CallByID(454152140, ...Array.prototype.slice.call(arguments, 0)); },

		/**
		 * App.GetTransalteMap
		 * GetTransalteMap 获取所有翻译配置
		 *
		 * @returns {Promise<string>}
		 **/
	    GetTransalteMap: function() { return wails.CallByID(3159850429, ...Array.prototype.slice.call(arguments, 0)); },

		/**
		 * App.GetTransalteWay
		 * GetTransalteWay 获取当前翻译的服务
		 *
		 * @returns {Promise<string>}
		 **/
	    GetTransalteWay: function() { return wails.CallByID(3877427008, ...Array.prototype.slice.call(arguments, 0)); },

		/**
		 * App.Hide
		 * Hide 通过名字控制窗口事件
		 * @param windowName {string}
		 * @returns {Promise<void>}
		 **/
	    Hide: function(windowName) { return wails.CallByID(3538035797, ...Array.prototype.slice.call(arguments, 0)); },

		/**
		 * App.MyFetch
		 * MyFetch URl
		 * @param URL {string}
 * @param content {map}
		 * @returns {Promise<>}
		 **/
	    MyFetch: function(URL, content) { return wails.CallByID(2071126117, ...Array.prototype.slice.call(arguments, 0)); },

		/**
		 * App.SetTransalteWay
		 * SetTransalteWay 设置当前翻译服务
		 * @param translateWay {string}
		 * @returns {Promise<void>}
		 **/
	    SetTransalteWay: function(translateWay) { return wails.CallByID(1606326012, ...Array.prototype.slice.call(arguments, 0)); },

		/**
		 * App.Show
		 * Show 通过名字控制窗口事件
		 * @param windowName {string}
		 * @returns {Promise<void>}
		 **/
	    Show: function(windowName) { return wails.CallByID(2781088484, ...Array.prototype.slice.call(arguments, 0)); },

		/**
		 * App.ToolBarShow
		 * ToolBarShow 显示工具弹窗，控制大小，布局, 前端调用，传递文本高度
		 * @param height {number}
		 * @returns {Promise<void>}
		 **/
	    ToolBarShow: function(height) { return wails.CallByID(2532748761, ...Array.prototype.slice.call(arguments, 0)); },

		/**
		 * App.Transalte
		 * Transalte 翻译逻辑
		 * @param queryText {string}
 * @param fromLang {string}
 * @param toLang {string}
		 * @returns {Promise<string>}
		 **/
	    Transalte: function(queryText, fromLang, toLang) { return wails.CallByID(3553729015, ...Array.prototype.slice.call(arguments, 0)); },
    },
};
