#!/usr/bin/env bash
# sync.sh — обновить static/vibe.html из основного стенда и пересобрать бинарь.
# Автоматизирует описанный в README порядок: cp, вычёркивание ссылки
# «назад к тестам», пересборка. Запускать из корня репозитория или так:
#   ./sync.sh
set -euo pipefail
cd "$(dirname "$0")"

SRC=../internal/webui/static/vibe.html
[[ -f "$SRC" ]] || { echo "нет исходника: $SRC — ожидаемый путь относительно vibe-only/" >&2; exit 1; }

cp "$SRC" static/vibe.html
# в основном стенде ссылка ведёт на страницу тестов; здесь страницы тестов нет — убираем
sed -i '/← назад к тестам/d' static/vibe.html
# здесь это не «SIP Tester», а самостоятельная страница — свой заголовок вкладки
sed -i 's|<title>.*</title>|<title>🎧 vibe mode</title>|' static/vibe.html

go vet ./...
go build -o vibe-only .
echo "готово: ./vibe-only"
