/** @type {import('tailwindcss').Config} */
module.exports = {
    mode: "all",
    content: [
        // include all rust, html and css files in the src directory
        "./src/**/*.{rs,html,css}",
        // include all html files in the output (dist) directory
        "./dist/**/*.html",
    ],
    theme: {
        fontSize: {
            sm: '7vmin',
            base: '10vmin',
            lg: '15vmin',
            xl: '40vmin',
        },
        extend: {
            colors: {
                primary: '#4050b5',
                secondary: '#3bb371',
            },
            fontFamily: {
                dice: ["dice", "monospace"],
            }
        },
    },
    plugins: [],
}
