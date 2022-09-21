# FE

C3 Frontend for Svelte

# Installing Svelte

npx degit sveltejs/template my-svelte-project
cd my-svelte-project

npm install
npm run dev

# Changing Port to 4000 for Frontend

1. Go to package.json file
1. Look for "scripts"
1. Add a `--port` flag: `"dev": "rollup -c -w --port 3000"`