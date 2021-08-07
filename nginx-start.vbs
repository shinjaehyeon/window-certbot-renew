Set WinScriptHost = CreateObject( "WScript.shell" )
PathName = CreateObject("Scripting.FileSystemObject").GetParentFolderName(WScript.ScriptFullName)
WinScriptHost.Run Chr(34) & PathName & "\nginx-start.bat" & Chr(34), 0 
Set WinScriptHost = Nothing
