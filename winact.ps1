Function GatherBoot {
    $bootTs = (Get-CimInstance -ClassName Win32_OperatingSystem).LastBootUpTime.ToFileTimeUTC() / 10000 - 11644473600000
    $ts = [int64](([datetime]::UtcNow)-(get-date "1/1/1970")).TotalMilliseconds
    Write-Host "winact boot=$bootTs $ts"
}

$GatherLock = {
    $ts = [int64](([datetime]::UtcNow)-(get-date "1/1/1970")).TotalMilliseconds
    Write-Host "winact locked=true $ts"
}

$GatherUnlock = {
    $ts = [int64](([datetime]::UtcNow)-(get-date "1/1/1970")).TotalMilliseconds
    Write-Host "winact locked=false $ts"
}

Function SetupLockEvents {
    $lockedIdentifier = "Locked"
    $lockedQuery = "Select * from win32_ProcessStartTrace where processname = 'logonui.exe'"
    Unregister-Event -SourceIdentifier $lockedIdentifier -ErrorAction SilentlyContinue
    Register-CimIndicationEvent -Query $lockedQuery -SourceIdentifier $lockedIdentifier -Action $GatherLock
    
    $unlockedIdentifier = "Unlocked"
    $unlockedQuery = "Select * from win32_ProcessStopTrace where processname = 'logonui.exe'"
    Unregister-Event -SourceIdentifier $unlockedIdentifier -ErrorAction SilentlyContinue
    Register-CimIndicationEvent -Query $unlockedQuery -SourceIdentifier $unlockedIdentifier -Action $GatherUnlock
}

SetupLockEvents

GatherBoot
