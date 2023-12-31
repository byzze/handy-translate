package events

type ApplicationEventType uint
type WindowEventType uint

var Common = newCommonEvents()

type commonEvents struct {
	ApplicationStarted ApplicationEventType
	WindowMaximise     WindowEventType
	WindowUnMaximise   WindowEventType
	WindowFullscreen   WindowEventType
	WindowUnFullscreen WindowEventType
	WindowRestore      WindowEventType
	WindowMinimise     WindowEventType
	WindowUnMinimise   WindowEventType
	WindowClosing      WindowEventType
	WindowZoom         WindowEventType
	WindowZoomIn       WindowEventType
	WindowZoomOut      WindowEventType
	WindowZoomReset    WindowEventType
	WindowFocus        WindowEventType
	WindowLostFocus    WindowEventType
	WindowShow         WindowEventType
	WindowHide         WindowEventType
	WindowDPIChanged   WindowEventType
	WindowFilesDropped WindowEventType
	ThemeChanged       ApplicationEventType
}

func newCommonEvents() commonEvents {
	return commonEvents{
		ApplicationStarted: 1173,
		WindowMaximise:     1174,
		WindowUnMaximise:   1175,
		WindowFullscreen:   1176,
		WindowUnFullscreen: 1177,
		WindowRestore:      1178,
		WindowMinimise:     1179,
		WindowUnMinimise:   1180,
		WindowClosing:      1181,
		WindowZoom:         1182,
		WindowZoomIn:       1183,
		WindowZoomOut:      1184,
		WindowZoomReset:    1185,
		WindowFocus:        1186,
		WindowLostFocus:    1187,
		WindowShow:         1188,
		WindowHide:         1189,
		WindowDPIChanged:   1190,
		WindowFilesDropped: 1191,
		ThemeChanged:       1192,
	}
}

var Linux = newLinuxEvents()

type linuxEvents struct {
	SystemThemeChanged ApplicationEventType
}

func newLinuxEvents() linuxEvents {
	return linuxEvents{
		SystemThemeChanged: 1193,
	}
}

var Mac = newMacEvents()

type macEvents struct {
	ApplicationDidBecomeActive                              ApplicationEventType
	ApplicationDidChangeBackingProperties                   ApplicationEventType
	ApplicationDidChangeEffectiveAppearance                 ApplicationEventType
	ApplicationDidChangeIcon                                ApplicationEventType
	ApplicationDidChangeOcclusionState                      ApplicationEventType
	ApplicationDidChangeScreenParameters                    ApplicationEventType
	ApplicationDidChangeStatusBarFrame                      ApplicationEventType
	ApplicationDidChangeStatusBarOrientation                ApplicationEventType
	ApplicationDidFinishLaunching                           ApplicationEventType
	ApplicationDidHide                                      ApplicationEventType
	ApplicationDidResignActiveNotification                  ApplicationEventType
	ApplicationDidUnhide                                    ApplicationEventType
	ApplicationDidUpdate                                    ApplicationEventType
	ApplicationWillBecomeActive                             ApplicationEventType
	ApplicationWillFinishLaunching                          ApplicationEventType
	ApplicationWillHide                                     ApplicationEventType
	ApplicationWillResignActive                             ApplicationEventType
	ApplicationWillTerminate                                ApplicationEventType
	ApplicationWillUnhide                                   ApplicationEventType
	ApplicationWillUpdate                                   ApplicationEventType
	ApplicationDidChangeTheme                               ApplicationEventType
	ApplicationShouldHandleReopen                           ApplicationEventType
	WindowDidBecomeKey                                      WindowEventType
	WindowDidBecomeMain                                     WindowEventType
	WindowDidBeginSheet                                     WindowEventType
	WindowDidChangeAlpha                                    WindowEventType
	WindowDidChangeBackingLocation                          WindowEventType
	WindowDidChangeBackingProperties                        WindowEventType
	WindowDidChangeCollectionBehavior                       WindowEventType
	WindowDidChangeEffectiveAppearance                      WindowEventType
	WindowDidChangeOcclusionState                           WindowEventType
	WindowDidChangeOrderingMode                             WindowEventType
	WindowDidChangeScreen                                   WindowEventType
	WindowDidChangeScreenParameters                         WindowEventType
	WindowDidChangeScreenProfile                            WindowEventType
	WindowDidChangeScreenSpace                              WindowEventType
	WindowDidChangeScreenSpaceProperties                    WindowEventType
	WindowDidChangeSharingType                              WindowEventType
	WindowDidChangeSpace                                    WindowEventType
	WindowDidChangeSpaceOrderingMode                        WindowEventType
	WindowDidChangeTitle                                    WindowEventType
	WindowDidChangeToolbar                                  WindowEventType
	WindowDidChangeVisibility                               WindowEventType
	WindowDidDeminiaturize                                  WindowEventType
	WindowDidEndSheet                                       WindowEventType
	WindowDidEnterFullScreen                                WindowEventType
	WindowDidEnterVersionBrowser                            WindowEventType
	WindowDidExitFullScreen                                 WindowEventType
	WindowDidExitVersionBrowser                             WindowEventType
	WindowDidExpose                                         WindowEventType
	WindowDidFocus                                          WindowEventType
	WindowDidMiniaturize                                    WindowEventType
	WindowDidMove                                           WindowEventType
	WindowDidOrderOffScreen                                 WindowEventType
	WindowDidOrderOnScreen                                  WindowEventType
	WindowDidResignKey                                      WindowEventType
	WindowDidResignMain                                     WindowEventType
	WindowDidResize                                         WindowEventType
	WindowDidUpdate                                         WindowEventType
	WindowDidUpdateAlpha                                    WindowEventType
	WindowDidUpdateCollectionBehavior                       WindowEventType
	WindowDidUpdateCollectionProperties                     WindowEventType
	WindowDidUpdateShadow                                   WindowEventType
	WindowDidUpdateTitle                                    WindowEventType
	WindowDidUpdateToolbar                                  WindowEventType
	WindowDidUpdateVisibility                               WindowEventType
	WindowShouldClose                                       WindowEventType
	WindowWillBecomeKey                                     WindowEventType
	WindowWillBecomeMain                                    WindowEventType
	WindowWillBeginSheet                                    WindowEventType
	WindowWillChangeOrderingMode                            WindowEventType
	WindowWillClose                                         WindowEventType
	WindowWillDeminiaturize                                 WindowEventType
	WindowWillEnterFullScreen                               WindowEventType
	WindowWillEnterVersionBrowser                           WindowEventType
	WindowWillExitFullScreen                                WindowEventType
	WindowWillExitVersionBrowser                            WindowEventType
	WindowWillFocus                                         WindowEventType
	WindowWillMiniaturize                                   WindowEventType
	WindowWillMove                                          WindowEventType
	WindowWillOrderOffScreen                                WindowEventType
	WindowWillOrderOnScreen                                 WindowEventType
	WindowWillResignMain                                    WindowEventType
	WindowWillResize                                        WindowEventType
	WindowWillUnfocus                                       WindowEventType
	WindowWillUpdate                                        WindowEventType
	WindowWillUpdateAlpha                                   WindowEventType
	WindowWillUpdateCollectionBehavior                      WindowEventType
	WindowWillUpdateCollectionProperties                    WindowEventType
	WindowWillUpdateShadow                                  WindowEventType
	WindowWillUpdateTitle                                   WindowEventType
	WindowWillUpdateToolbar                                 WindowEventType
	WindowWillUpdateVisibility                              WindowEventType
	WindowWillUseStandardFrame                              WindowEventType
	MenuWillOpen                                            ApplicationEventType
	MenuDidOpen                                             ApplicationEventType
	MenuDidClose                                            ApplicationEventType
	MenuWillSendAction                                      ApplicationEventType
	MenuDidSendAction                                       ApplicationEventType
	MenuWillHighlightItem                                   ApplicationEventType
	MenuDidHighlightItem                                    ApplicationEventType
	MenuWillDisplayItem                                     ApplicationEventType
	MenuDidDisplayItem                                      ApplicationEventType
	MenuWillAddItem                                         ApplicationEventType
	MenuDidAddItem                                          ApplicationEventType
	MenuWillRemoveItem                                      ApplicationEventType
	MenuDidRemoveItem                                       ApplicationEventType
	MenuWillBeginTracking                                   ApplicationEventType
	MenuDidBeginTracking                                    ApplicationEventType
	MenuWillEndTracking                                     ApplicationEventType
	MenuDidEndTracking                                      ApplicationEventType
	MenuWillUpdate                                          ApplicationEventType
	MenuDidUpdate                                           ApplicationEventType
	MenuWillPopUp                                           ApplicationEventType
	MenuDidPopUp                                            ApplicationEventType
	MenuWillSendActionToItem                                ApplicationEventType
	MenuDidSendActionToItem                                 ApplicationEventType
	WebViewDidStartProvisionalNavigation                    WindowEventType
	WebViewDidReceiveServerRedirectForProvisionalNavigation WindowEventType
	WebViewDidFinishNavigation                              WindowEventType
	WebViewDidCommitNavigation                              WindowEventType
	WindowFileDraggingEntered                               WindowEventType
	WindowFileDraggingPerformed                             WindowEventType
	WindowFileDraggingExited                                WindowEventType
}

func newMacEvents() macEvents {
	return macEvents{
		ApplicationDidBecomeActive:               1024,
		ApplicationDidChangeBackingProperties:    1025,
		ApplicationDidChangeEffectiveAppearance:  1026,
		ApplicationDidChangeIcon:                 1027,
		ApplicationDidChangeOcclusionState:       1028,
		ApplicationDidChangeScreenParameters:     1029,
		ApplicationDidChangeStatusBarFrame:       1030,
		ApplicationDidChangeStatusBarOrientation: 1031,
		ApplicationDidFinishLaunching:            1032,
		ApplicationDidHide:                       1033,
		ApplicationDidResignActiveNotification:   1034,
		ApplicationDidUnhide:                     1035,
		ApplicationDidUpdate:                     1036,
		ApplicationWillBecomeActive:              1037,
		ApplicationWillFinishLaunching:           1038,
		ApplicationWillHide:                      1039,
		ApplicationWillResignActive:              1040,
		ApplicationWillTerminate:                 1041,
		ApplicationWillUnhide:                    1042,
		ApplicationWillUpdate:                    1043,
		ApplicationDidChangeTheme:                1044,
		ApplicationShouldHandleReopen:            1045,
		WindowDidBecomeKey:                       1046,
		WindowDidBecomeMain:                      1047,
		WindowDidBeginSheet:                      1048,
		WindowDidChangeAlpha:                     1049,
		WindowDidChangeBackingLocation:           1050,
		WindowDidChangeBackingProperties:         1051,
		WindowDidChangeCollectionBehavior:        1052,
		WindowDidChangeEffectiveAppearance:       1053,
		WindowDidChangeOcclusionState:            1054,
		WindowDidChangeOrderingMode:              1055,
		WindowDidChangeScreen:                    1056,
		WindowDidChangeScreenParameters:          1057,
		WindowDidChangeScreenProfile:             1058,
		WindowDidChangeScreenSpace:               1059,
		WindowDidChangeScreenSpaceProperties:     1060,
		WindowDidChangeSharingType:               1061,
		WindowDidChangeSpace:                     1062,
		WindowDidChangeSpaceOrderingMode:         1063,
		WindowDidChangeTitle:                     1064,
		WindowDidChangeToolbar:                   1065,
		WindowDidChangeVisibility:                1066,
		WindowDidDeminiaturize:                   1067,
		WindowDidEndSheet:                        1068,
		WindowDidEnterFullScreen:                 1069,
		WindowDidEnterVersionBrowser:             1070,
		WindowDidExitFullScreen:                  1071,
		WindowDidExitVersionBrowser:              1072,
		WindowDidExpose:                          1073,
		WindowDidFocus:                           1074,
		WindowDidMiniaturize:                     1075,
		WindowDidMove:                            1076,
		WindowDidOrderOffScreen:                  1077,
		WindowDidOrderOnScreen:                   1078,
		WindowDidResignKey:                       1079,
		WindowDidResignMain:                      1080,
		WindowDidResize:                          1081,
		WindowDidUpdate:                          1082,
		WindowDidUpdateAlpha:                     1083,
		WindowDidUpdateCollectionBehavior:        1084,
		WindowDidUpdateCollectionProperties:      1085,
		WindowDidUpdateShadow:                    1086,
		WindowDidUpdateTitle:                     1087,
		WindowDidUpdateToolbar:                   1088,
		WindowDidUpdateVisibility:                1089,
		WindowShouldClose:                        1090,
		WindowWillBecomeKey:                      1091,
		WindowWillBecomeMain:                     1092,
		WindowWillBeginSheet:                     1093,
		WindowWillChangeOrderingMode:             1094,
		WindowWillClose:                          1095,
		WindowWillDeminiaturize:                  1096,
		WindowWillEnterFullScreen:                1097,
		WindowWillEnterVersionBrowser:            1098,
		WindowWillExitFullScreen:                 1099,
		WindowWillExitVersionBrowser:             1100,
		WindowWillFocus:                          1101,
		WindowWillMiniaturize:                    1102,
		WindowWillMove:                           1103,
		WindowWillOrderOffScreen:                 1104,
		WindowWillOrderOnScreen:                  1105,
		WindowWillResignMain:                     1106,
		WindowWillResize:                         1107,
		WindowWillUnfocus:                        1108,
		WindowWillUpdate:                         1109,
		WindowWillUpdateAlpha:                    1110,
		WindowWillUpdateCollectionBehavior:       1111,
		WindowWillUpdateCollectionProperties:     1112,
		WindowWillUpdateShadow:                   1113,
		WindowWillUpdateTitle:                    1114,
		WindowWillUpdateToolbar:                  1115,
		WindowWillUpdateVisibility:               1116,
		WindowWillUseStandardFrame:               1117,
		MenuWillOpen:                             1118,
		MenuDidOpen:                              1119,
		MenuDidClose:                             1120,
		MenuWillSendAction:                       1121,
		MenuDidSendAction:                        1122,
		MenuWillHighlightItem:                    1123,
		MenuDidHighlightItem:                     1124,
		MenuWillDisplayItem:                      1125,
		MenuDidDisplayItem:                       1126,
		MenuWillAddItem:                          1127,
		MenuDidAddItem:                           1128,
		MenuWillRemoveItem:                       1129,
		MenuDidRemoveItem:                        1130,
		MenuWillBeginTracking:                    1131,
		MenuDidBeginTracking:                     1132,
		MenuWillEndTracking:                      1133,
		MenuDidEndTracking:                       1134,
		MenuWillUpdate:                           1135,
		MenuDidUpdate:                            1136,
		MenuWillPopUp:                            1137,
		MenuDidPopUp:                             1138,
		MenuWillSendActionToItem:                 1139,
		MenuDidSendActionToItem:                  1140,
		WebViewDidStartProvisionalNavigation:     1141,
		WebViewDidReceiveServerRedirectForProvisionalNavigation: 1142,
		WebViewDidFinishNavigation:                              1143,
		WebViewDidCommitNavigation:                              1144,
		WindowFileDraggingEntered:                               1145,
		WindowFileDraggingPerformed:                             1146,
		WindowFileDraggingExited:                                1147,
	}
}

var Windows = newWindowsEvents()

type windowsEvents struct {
	SystemThemeChanged         ApplicationEventType
	APMPowerStatusChange       ApplicationEventType
	APMSuspend                 ApplicationEventType
	APMResumeAutomatic         ApplicationEventType
	APMResumeSuspend           ApplicationEventType
	APMPowerSettingChange      ApplicationEventType
	ApplicationStarted         ApplicationEventType
	WebViewNavigationCompleted WindowEventType
	WindowInactive             WindowEventType
	WindowActive               WindowEventType
	WindowClickActive          WindowEventType
	WindowMaximise             WindowEventType
	WindowUnMaximise           WindowEventType
	WindowFullscreen           WindowEventType
	WindowUnFullscreen         WindowEventType
	WindowRestore              WindowEventType
	WindowMinimise             WindowEventType
	WindowUnMinimise           WindowEventType
	WindowClose                WindowEventType
	WindowSetFocus             WindowEventType
	WindowKillFocus            WindowEventType
	WindowDragDrop             WindowEventType
	WindowDragEnter            WindowEventType
	WindowDragLeave            WindowEventType
	WindowDragOver             WindowEventType
}

func newWindowsEvents() windowsEvents {
	return windowsEvents{
		SystemThemeChanged:         1148,
		APMPowerStatusChange:       1149,
		APMSuspend:                 1150,
		APMResumeAutomatic:         1151,
		APMResumeSuspend:           1152,
		APMPowerSettingChange:      1153,
		ApplicationStarted:         1154,
		WebViewNavigationCompleted: 1155,
		WindowInactive:             1156,
		WindowActive:               1157,
		WindowClickActive:          1158,
		WindowMaximise:             1159,
		WindowUnMaximise:           1160,
		WindowFullscreen:           1161,
		WindowUnFullscreen:         1162,
		WindowRestore:              1163,
		WindowMinimise:             1164,
		WindowUnMinimise:           1165,
		WindowClose:                1166,
		WindowSetFocus:             1167,
		WindowKillFocus:            1168,
		WindowDragDrop:             1169,
		WindowDragEnter:            1170,
		WindowDragLeave:            1171,
		WindowDragOver:             1172,
	}
}

func JSEvent(event uint) string {
	return eventToJS[event]
}

var eventToJS = map[uint]string{
	1024: "mac:ApplicationDidBecomeActive",
	1025: "mac:ApplicationDidChangeBackingProperties",
	1026: "mac:ApplicationDidChangeEffectiveAppearance",
	1027: "mac:ApplicationDidChangeIcon",
	1028: "mac:ApplicationDidChangeOcclusionState",
	1029: "mac:ApplicationDidChangeScreenParameters",
	1030: "mac:ApplicationDidChangeStatusBarFrame",
	1031: "mac:ApplicationDidChangeStatusBarOrientation",
	1032: "mac:ApplicationDidFinishLaunching",
	1033: "mac:ApplicationDidHide",
	1034: "mac:ApplicationDidResignActiveNotification",
	1035: "mac:ApplicationDidUnhide",
	1036: "mac:ApplicationDidUpdate",
	1037: "mac:ApplicationWillBecomeActive",
	1038: "mac:ApplicationWillFinishLaunching",
	1039: "mac:ApplicationWillHide",
	1040: "mac:ApplicationWillResignActive",
	1041: "mac:ApplicationWillTerminate",
	1042: "mac:ApplicationWillUnhide",
	1043: "mac:ApplicationWillUpdate",
	1044: "mac:ApplicationDidChangeTheme!",
	1045: "mac:ApplicationShouldHandleReopen!",
	1046: "mac:WindowDidBecomeKey",
	1047: "mac:WindowDidBecomeMain",
	1048: "mac:WindowDidBeginSheet",
	1049: "mac:WindowDidChangeAlpha",
	1050: "mac:WindowDidChangeBackingLocation",
	1051: "mac:WindowDidChangeBackingProperties",
	1052: "mac:WindowDidChangeCollectionBehavior",
	1053: "mac:WindowDidChangeEffectiveAppearance",
	1054: "mac:WindowDidChangeOcclusionState",
	1055: "mac:WindowDidChangeOrderingMode",
	1056: "mac:WindowDidChangeScreen",
	1057: "mac:WindowDidChangeScreenParameters",
	1058: "mac:WindowDidChangeScreenProfile",
	1059: "mac:WindowDidChangeScreenSpace",
	1060: "mac:WindowDidChangeScreenSpaceProperties",
	1061: "mac:WindowDidChangeSharingType",
	1062: "mac:WindowDidChangeSpace",
	1063: "mac:WindowDidChangeSpaceOrderingMode",
	1064: "mac:WindowDidChangeTitle",
	1065: "mac:WindowDidChangeToolbar",
	1066: "mac:WindowDidChangeVisibility",
	1067: "mac:WindowDidDeminiaturize",
	1068: "mac:WindowDidEndSheet",
	1069: "mac:WindowDidEnterFullScreen",
	1070: "mac:WindowDidEnterVersionBrowser",
	1071: "mac:WindowDidExitFullScreen",
	1072: "mac:WindowDidExitVersionBrowser",
	1073: "mac:WindowDidExpose",
	1074: "mac:WindowDidFocus",
	1075: "mac:WindowDidMiniaturize",
	1076: "mac:WindowDidMove",
	1077: "mac:WindowDidOrderOffScreen",
	1078: "mac:WindowDidOrderOnScreen",
	1079: "mac:WindowDidResignKey",
	1080: "mac:WindowDidResignMain",
	1081: "mac:WindowDidResize",
	1082: "mac:WindowDidUpdate",
	1083: "mac:WindowDidUpdateAlpha",
	1084: "mac:WindowDidUpdateCollectionBehavior",
	1085: "mac:WindowDidUpdateCollectionProperties",
	1086: "mac:WindowDidUpdateShadow",
	1087: "mac:WindowDidUpdateTitle",
	1088: "mac:WindowDidUpdateToolbar",
	1089: "mac:WindowDidUpdateVisibility",
	1090: "mac:WindowShouldClose!",
	1091: "mac:WindowWillBecomeKey",
	1092: "mac:WindowWillBecomeMain",
	1093: "mac:WindowWillBeginSheet",
	1094: "mac:WindowWillChangeOrderingMode",
	1095: "mac:WindowWillClose",
	1096: "mac:WindowWillDeminiaturize",
	1097: "mac:WindowWillEnterFullScreen",
	1098: "mac:WindowWillEnterVersionBrowser",
	1099: "mac:WindowWillExitFullScreen",
	1100: "mac:WindowWillExitVersionBrowser",
	1101: "mac:WindowWillFocus",
	1102: "mac:WindowWillMiniaturize",
	1103: "mac:WindowWillMove",
	1104: "mac:WindowWillOrderOffScreen",
	1105: "mac:WindowWillOrderOnScreen",
	1106: "mac:WindowWillResignMain",
	1107: "mac:WindowWillResize",
	1108: "mac:WindowWillUnfocus",
	1109: "mac:WindowWillUpdate",
	1110: "mac:WindowWillUpdateAlpha",
	1111: "mac:WindowWillUpdateCollectionBehavior",
	1112: "mac:WindowWillUpdateCollectionProperties",
	1113: "mac:WindowWillUpdateShadow",
	1114: "mac:WindowWillUpdateTitle",
	1115: "mac:WindowWillUpdateToolbar",
	1116: "mac:WindowWillUpdateVisibility",
	1117: "mac:WindowWillUseStandardFrame",
	1118: "mac:MenuWillOpen",
	1119: "mac:MenuDidOpen",
	1120: "mac:MenuDidClose",
	1121: "mac:MenuWillSendAction",
	1122: "mac:MenuDidSendAction",
	1123: "mac:MenuWillHighlightItem",
	1124: "mac:MenuDidHighlightItem",
	1125: "mac:MenuWillDisplayItem",
	1126: "mac:MenuDidDisplayItem",
	1127: "mac:MenuWillAddItem",
	1128: "mac:MenuDidAddItem",
	1129: "mac:MenuWillRemoveItem",
	1130: "mac:MenuDidRemoveItem",
	1131: "mac:MenuWillBeginTracking",
	1132: "mac:MenuDidBeginTracking",
	1133: "mac:MenuWillEndTracking",
	1134: "mac:MenuDidEndTracking",
	1135: "mac:MenuWillUpdate",
	1136: "mac:MenuDidUpdate",
	1137: "mac:MenuWillPopUp",
	1138: "mac:MenuDidPopUp",
	1139: "mac:MenuWillSendActionToItem",
	1140: "mac:MenuDidSendActionToItem",
	1141: "mac:WebViewDidStartProvisionalNavigation",
	1142: "mac:WebViewDidReceiveServerRedirectForProvisionalNavigation",
	1143: "mac:WebViewDidFinishNavigation",
	1144: "mac:WebViewDidCommitNavigation",
	1145: "mac:WindowFileDraggingEntered",
	1146: "mac:WindowFileDraggingPerformed",
	1147: "mac:WindowFileDraggingExited",
	1148: "windows:SystemThemeChanged",
	1149: "windows:APMPowerStatusChange",
	1150: "windows:APMSuspend",
	1151: "windows:APMResumeAutomatic",
	1152: "windows:APMResumeSuspend",
	1153: "windows:APMPowerSettingChange",
	1154: "windows:ApplicationStarted",
	1155: "windows:WebViewNavigationCompleted",
	1156: "windows:WindowInactive",
	1157: "windows:WindowActive",
	1158: "windows:WindowClickActive",
	1159: "windows:WindowMaximise",
	1160: "windows:WindowUnMaximise",
	1161: "windows:WindowFullscreen",
	1162: "windows:WindowUnFullscreen",
	1163: "windows:WindowRestore",
	1164: "windows:WindowMinimise",
	1165: "windows:WindowUnMinimise",
	1166: "windows:WindowClose",
	1167: "windows:WindowSetFocus",
	1168: "windows:WindowKillFocus",
	1169: "windows:WindowDragDrop",
	1170: "windows:WindowDragEnter",
	1171: "windows:WindowDragLeave",
	1172: "windows:WindowDragOver",
	1173: "common:ApplicationStarted",
	1174: "common:WindowMaximise",
	1175: "common:WindowUnMaximise",
	1176: "common:WindowFullscreen",
	1177: "common:WindowUnFullscreen",
	1178: "common:WindowRestore",
	1179: "common:WindowMinimise",
	1180: "common:WindowUnMinimise",
	1181: "common:WindowClosing",
	1182: "common:WindowZoom",
	1183: "common:WindowZoomIn",
	1184: "common:WindowZoomOut",
	1185: "common:WindowZoomReset",
	1186: "common:WindowFocus",
	1187: "common:WindowLostFocus",
	1188: "common:WindowShow",
	1189: "common:WindowHide",
	1190: "common:WindowDPIChanged",
	1191: "common:WindowFilesDropped",
	1192: "common:ThemeChanged",
	1193: "linux:SystemThemeChanged",
}
