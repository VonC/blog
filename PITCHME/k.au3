#include <WinAPI.au3>
#Include <WindowsConstants.au3>

Global $sHexKeys, $sMouse, $sString, $hHookKeyboard, $pStub_KeyProc

HotKeySet("{PAUSE}", "PauseNow");
HotKeySet("^{BROWSER_HOME}", "VK_BROWSER_HOME")

$pStub_KeyProc = DllCallbackRegister("_KeyProc", "int", "int;ptr;ptr")
$hHookKeyboard = _WinAPI_SetWindowsHookEx($WH_KEYBOARD_LL, DllCallbackGetPtr($pStub_KeyProc), _WinAPI_GetModuleHandle(0), 0)

While 1
    Sleep(10)
WEnd

Func PauseNow()
    WinMinimizeAll()
EndFunc  ;==>ExitNow
Func OnAutoITExit()
    DllCallbackFree($pStub_KeyProc)
    _WinAPI_UnhookWindowsHookEx($hHookKeyboard)
EndFunc  ;==>OnAutoITExit

Func _KeyProc($nCode, $wParam, $lParam)
    If $nCode < 0 Then Return _WinAPI_CallNextHookEx($hHookKeyboard, $nCode, $wParam, $lParam)
    Local $KBDLLHOOKSTRUCT = DllStructCreate("dword vkCode;dword scanCode;dword flags;dword time;ptr dwExtraInfo", $lParam)
    Local $flag = DllStructGetData($KBDLLHOOKSTRUCT, "flag")
    Local $vkCode = DllStructGetData($KBDLLHOOKSTRUCT, "vkCode")
    Local $dwExtraInfo = DllStructGetData($KBDLLHOOKSTRUCT, "dwExtraInfo")
	Local $sText = WinGetTitle("[active]")
	If _StringStartsWith($sText, '[ GitPitch ]') <= 0 Then Return _WinAPI_CallNextHookEx($hHookKeyboard, $nCode, $wParam, $lParam)
	Switch $wParam
    Case $WM_KEYDOWN, $WM_SYSKEYDOWN
       ConsoleWrite("$WM_KEYDOWN: flag - " & DllStructGetData($KBDLLHOOKSTRUCT, "flag") & @TAB & "vkCode - " & DllStructGetData($KBDLLHOOKSTRUCT, "vkCode") & @TAB & "dwExtraInfo - " & $dwExtraInfo & @CRLF)
	   Switch $vkCode
			Case 0x41 ;KeyPressed========= 0x41 = "a"
			    _keybd_event(0x60, 0) ;Key to Send to OS=====0x60 = numpad0
                Return -1
			Case 0x26 ;KeyPressed========= 0x26 = "UP ARROW key"
			    _keybd_event(0x20, 0) ;Key to Send to OS=====0x20 = SPACEBAR
                Return -1
			Case 0x28 ;KeyPressed========= 0x28 = "DOWN ARROW key"
			    _keybd_event(0x25, 0) ;Key to Send to OS=====0x25 = LEFT ARROW key
                Return -1
			Case 0xAE ;KeyPressed========= 0xAE = "Volume Down key"
			    _keybd_event(0x28 , 0) ;Key to Send to OS=====0x28 = DOWN ARROW key
                Return -1
			Case 0xAF ;KeyPressed========= 0xAF = "Volume Up key"
			    _keybd_event(0x26, 0) ;Key to Send to OS=====0x26 = UP ARROW key²
                Return -1
			Case 0x09 ;KeyPressed========= 0x09 = "TAB key"
			    _keybd_event(0x46, 0) ;Key to Send to OS=====0x46 = F key
                Return -1
			Case 0xA4 ;KeyPressed========= 0xA4 = "Left Menu key"
			    _keybd_event(0x42, 0) ;Key to Send to OS=====0x42 = B key
                Return -1

			 Case 172;KeyPressed========= 172 = "VK_BROWSER_HOME"
			   _keybd_event(112, 0) ;Key to Send to OS=====112 = F1 key
			    Return -1
			 Case 166;KeyPressed========= 166 = "VK_BROWSER_BACK"
				If $scanCode <> 29 Then
			    _keybd_event(27, 0) ;Key to Send to OS=====27 = 	ESC key
				Else
				  _keybd_event(166, 0) ;Key to Send to OS=====112 = F1 key
				EndIf
                Return -1

	  EndSwitch
   EndSwitch
    Return _WinAPI_CallNextHookEx($hHookKeyboard, $nCode, $wParam, $lParam)
EndFunc  ;==>_KeyProc

Func VK_BROWSER_HOME()
   _keybd_event(172, 0)
EndFunc  ;==>ExitNow

Func _keybd_event($vkCode, $Flag)
    DllCall('user32.dll', 'int', 'keybd_event', 'int', $vkCode, 'int', 0, 'int', $Flag, 'ptr', 0)
 EndFunc; _keybd_event

Func _StringStartsWith($string, $start, $case = 0)
    If StringLen($start) > StringLen($string) Then Return -1
    If $case > 0 Then
        If StringLeft($string, StringLen($start)) == $start Then Return 1
    Else
        If StringLeft($string, StringLen($start)) = $start Then Return 1
    EndIf
    Return 0
EndFunc   ;==>_StringStartsWith