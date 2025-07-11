
$ProjectName = "E-Commerce Microservice API"
$ComposeFile = "docker-compose.yml"
$Services = @("backend-api", "backend-worker", "directus", "n8n", "postgres", "redis", "prometheus", "grafana")

# --- Warna untuk Output ---
$ColorInfo = "Cyan"
$ColorSuccess = "Green"
$ColorWarning = "Yellow"
$ColorError = "Red"
$ColorTitle = "White"

function Test-Prerequisites {
    $dockerExists = Get-Command docker -ErrorAction SilentlyContinue
    $composeExists = Get-Command docker-compose -ErrorAction SilentlyContinue
    if (-not $dockerExists -or -not $composeExists) {
        Write-Host "ERROR: Docker dan Docker Compose tidak ditemukan." -ForegroundColor $ColorError
        Write-Host "Pastikan Docker Desktop sudah terinstal dan berjalan." -ForegroundColor $ColorError
        exit
    }
}

function Start-Project {
    Write-Host "Membangun dan memulai semua container di background..." -ForegroundColor $ColorInfo
    docker-compose -f $ComposeFile up --build -d
    if ($LASTEXITCODE -eq 0) {
        Write-Host "`nProject berhasil dimulai." -ForegroundColor $ColorSuccess
        Start-Sleep -Seconds 2
        Get-ProjectStatus
    } else {
        Write-Host "`nTerjadi kesalahan saat memulai project." -ForegroundColor $ColorError
    }
}

function Stop-Project {
    Write-Host "Menghentikan semua container project..." -ForegroundColor $ColorInfo
    docker-compose -f $ComposeFile down
    Write-Host "Semua container telah dihentikan." -ForegroundColor $ColorSuccess
}

function Get-ProjectStatus {
    Write-Host "Status container saat ini:" -ForegroundColor $ColorInfo
    docker-compose -f $ComposeFile ps
}

function View-Logs {
    while ($true) {
        Clear-Host
        Write-Host "--- Tampilkan Log Service ---" -ForegroundColor $ColorTitle
        for ($i = 0; $i -lt $Services.Length; $i++) {
            Write-Host (" " + ($i + 1) + ". " + $Services[$i])
        }
        Write-Host " 0. Kembali ke Menu Utama"
        $choice = Read-Host "`nPilih service untuk melihat log"

        if ($choice -eq '0') { break }

        $index = [int]$choice - 1
        if ($index -ge 0 -and $index -lt $Services.Length) {
            $serviceName = $Services[$index]
            Clear-Host
            Write-Host "Menampilkan log untuk '$serviceName'. Tekan CTRL+C untuk berhenti." -ForegroundColor $ColorInfo
            docker-compose -f $ComposeFile logs -f --tail="100" $serviceName
            Read-Host "Tekan ENTER untuk kembali ke menu log"
        }
    }
}

function Clean-Project {
    Write-Host "PERINGATAN: Aksi ini akan menghentikan semua container, menghapus network, dan MENGHAPUS SEMUA DATA di dalam volume (database, redis, n8n, dll)." -ForegroundColor $ColorWarning
    $confirmation = Read-Host "Apakah Anda benar-benar yakin? Ketik 'yes' untuk melanjutkan"
    if ($confirmation -eq 'yes') {
        Write-Host "Membersihkan project..." -ForegroundColor $ColorInfo
        docker-compose -f $ComposeFile down -v --remove-orphans
        if ($LASTEXITCODE -eq 0) {
            Write-Host "Project berhasil dibersihkan. Semua container, network, dan volume telah dihapus." -ForegroundColor $ColorSuccess
        } else {
            Write-Host "Terjadi kesalahan saat membersihkan project." -ForegroundColor $ColorError
        }
    } else {
        Write-Host "Aksi dibatalkan." -ForegroundColor $ColorInfo
    }
}

function Access-Shell {
    while ($true) {
        Clear-Host
        Write-Host "--- Akses Shell Container ---" -ForegroundColor $ColorTitle
        for ($i = 0; $i -lt $Services.Length; $i++) {
            Write-Host (" " + ($i + 1) + ". " + $Services[$i])
        }
        Write-Host " 0. Kembali ke Menu Utama"
        $choice = Read-Host "`nPilih service untuk diakses"

        if ($choice -eq '0') { break }

        $index = [int]$choice - 1
        if ($index -ge 0 -and $index -lt $Services.Length) {
            $serviceName = $Services[$index]
            Write-Host "Membuka shell untuk '$serviceName'. Ketik 'exit' untuk keluar." -ForegroundColor $ColorInfo
            docker-compose -f $ComposeFile exec $serviceName sh
        }
    }
}

function Restart-Project {
    Write-Host "Me-restart semua service..." -ForegroundColor $ColorInfo
    Stop-Project
    Start-Sleep -Seconds 2
    Start-Project
}


function Show-MainMenu {
    while ($true) {
        Clear-Host
        Write-Host "==========================================" -ForegroundColor $ColorTitle
        Write-Host "      $($ProjectName.ToUpper()) CLI"
        Write-Host "==========================================" -ForegroundColor $ColorTitle
        Write-Host ""
        Write-Host " 1. Start Project" -ForegroundColor $ColorInfo
        Write-Host " 2. Stop Project" -ForegroundColor $ColorInfo
        Write-Host " 3. Restart Project" -ForegroundColor $ColorInfo
        Write-Host " 4. Lihat Status" -ForegroundColor $ColorInfo
        Write-Host " 5. Tampilkan Log" -ForegroundColor $ColorInfo
        Write-Host " 6. Akses Shell Container" -ForegroundColor $ColorInfo
        Write-Host " 7. Clean Project (Hapus Semua & Data)" -ForegroundColor $ColorWarning
        Write-Host " Q. Keluar" -ForegroundColor $ColorInfo
        Write-Host ""
        $choice = Read-Host "Masukkan pilihan Anda"

        switch ($choice) {
            '1' { Start-Project }
            '2' { Stop-Project }
            '3' { Restart-Project }
            '4' { Get-ProjectStatus }
            '5' { View-Logs }
            '6' { Access-Shell }
            '7' { Clean-Project }
            'q' { exit }
            default { Write-Host "Pilihan tidak valid." -ForegroundColor $ColorError }
        }

        if ($choice -ne '5' -and $choice -ne '6') {
             Write-Host "`nTekan ENTER untuk melanjutkan..." -ForegroundColor $ColorInfo
             Read-Host | Out-Null
        }
    }
}

# --- Jalankan Script ---
Test-Prerequisites
Show-MainMenu