#!/bin/sh
set -e

if [ ! -d "node_modules" ]; then
    echo "Creating new Next.js project..."
    npx create-next-app@latest . --ts --tailwind --eslint --app --src-dir --use-npm --no-git --import-alias "@/*" --experimental-app --experimental-react 19 << ANSWERS
    yes
    yes
ANSWERS
fi

if [ -f "package.json" ] && [ ! -d "node_modules" ]; then
    npm install
fi

npm run dev