[CmdletBinding()]
param(
    [ValidateSet(
        'help',
        'up',
        'down',
        'run-api',
        'build-api',
        'install-ui1',
        'install-ui2',
        'run-ui1',
        'run-ui2',
        'build-ui1',
        'build-ui2',
        'migrate-up',
        'migrate-down',
        'db-seed',
        'tidy',
        'test'
    )]
    [string]$Task = 'help',

    [string]$AppName = 'gladiatortravel',
    [string]$DbUrl = 'postgres://postgres:postgres@localhost:5432/gladiatortravel?sslmode=disable',
    [int]$ApiPort = 8080,
    [int]$Ui1Port = 5173,
    [int]$Ui2Port = 5174
)

Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'

$RepoRoot = Split-Path -Parent $PSScriptRoot
Set-Location $RepoRoot

$npmCmdPath = Join-Path $env:ProgramFiles 'nodejs\npm.cmd'
$NpmCommand = if (Test-Path $npmCmdPath) { $npmCmdPath } else { 'npm' }

function Invoke-Step {
    param(
        [Parameter(Mandatory = $true)]
        [scriptblock]$Action
    )

    & $Action
    if ($LASTEXITCODE -ne 0) {
        throw "Command failed with exit code $LASTEXITCODE"
    }
}

function Show-Help {
    @"
Usage:
  .\scripts\dev.ps1 -Task <task> [options]

Tasks:
  help         Show this help
  up           Start docker services
  down         Stop docker services
  run-api      Run Go API
  build-api    Build Go API binary
  install-ui1  Install UI1 dependencies (web)
  install-ui2  Install UI2 dependencies (frontend)
  run-ui1      Run UI1 dev server (web)
  run-ui2      Run UI2 dev server (frontend)
  build-ui1    Build UI1 (web)
  build-ui2    Build UI2 (frontend)
  migrate-up   Run DB migrations up
  migrate-down Run DB migrations down by 1
  db-seed      Seed database
  tidy         Run go mod tidy
  test         Run go test ./...

Options:
  -ApiPort <int>   API port (default: 8080)
  -Ui1Port <int>   UI1 port (default: 5173)
  -Ui2Port <int>   UI2 port (default: 5174)
  -DbUrl <string>  Database URL
  -AppName <name>  API binary name (default: gladiatortravel)
"@
}

switch ($Task) {
    'help' {
        Show-Help
    }

    'up' {
        Invoke-Step { docker compose up -d }
    }

    'down' {
        Invoke-Step { docker compose down }
    }

    'run-api' {
        $env:PORT = "$ApiPort"
        $env:DATABASE_URL = $DbUrl
        Invoke-Step { go run ./cmd/api }
    }

    'build-api' {
        Invoke-Step { go build -o "$AppName.exe" ./cmd/api }
    }

    'install-ui1' {
        Invoke-Step { & $NpmCommand --prefix ./web install }
    }

    'install-ui2' {
        Invoke-Step { & $NpmCommand --prefix ./frontend install }
    }

    'run-ui1' {
        Invoke-Step { & $NpmCommand --prefix ./web run dev -- --port "$Ui1Port" }
    }

    'run-ui2' {
        Invoke-Step { & $NpmCommand --prefix ./frontend run dev -- --port "$Ui2Port" }
    }

    'build-ui1' {
        Invoke-Step { & $NpmCommand --prefix ./web run build }
    }

    'build-ui2' {
        Invoke-Step { & $NpmCommand --prefix ./frontend run build }
    }

    'migrate-up' {
        Invoke-Step { migrate -path migrations -database "$DbUrl" up }
    }

    'migrate-down' {
        Invoke-Step { migrate -path migrations -database "$DbUrl" down 1 }
    }

    'db-seed' {
        Invoke-Step { psql "$DbUrl" -f seed/001_base.sql }
        Invoke-Step { psql "$DbUrl" -f seed/002_scores_and_core_data.sql }
        Invoke-Step { psql "$DbUrl" -f seed/003_itinerary_samples.sql }
    }

    'tidy' {
        Invoke-Step { go mod tidy }
    }

    'test' {
        Invoke-Step { go test ./... }
    }
}
