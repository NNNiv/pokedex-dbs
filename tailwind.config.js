/** @type {import('tailwindcss').Config} */
// module.exports = {
//   content: ["./internal/templates/**/*.templ"],
//   theme: {
//     colors: {
//       pkmnbg: "#35374B",
//       pkmnfg: "#7DD696",
//       pkmnfg2: "#1B2729",
//       search: "#888DC6",
//       searchbtn: "#D9D9D9",
//       grass: "#008000",
//       fire: "#fd7d24",
//       water: "#0000FF",
//       electric: "#FFD700",
//       ice: "#00FFFF",
//       poison: "#800080",
//       psychic: "#FF1493",
//       bug: "#729f3f",
//       rock: "#A52A2A",
//       ghost: "#800080",
//       dragon: "#0000FF",
//       dark: "#707070",
//       steel: "#A9A9A9",
//       fairy: "#fdb9e9",
//       fighting: "#FF0000",
//       flying: "#0000FF",
//       ground: "#A52A2A",
//       normal: "#FFFFFF",
//     },
//     fontFamily: {
//       pkmn: ["PKMN RBYGSC"],
//     },
//   },
// };

// bug - #729f3f
// dark - #707070
// fairy - #fdb9e9
// fire - #fd7d24
// ghost - #7b62a3
// ground - #ab9842
// normal - #a4acaf
// psychic - #f366b9
// steel - #9eb7b8
// electric - #9eb7b8
// fighting - #d56723
// flying - #3dc7ef
// grass - #9bcc50
// ice - #51c4e7
// poison - #b97fc9
// rock - #b97fc9
// water - #4592c4

// rewrite the tailwindcss config to use the above colors
// and the PKMN RBYGSC font
// make sure to only use the colors listed above and the font listed above

module.exports = {
  content: ["./internal/templates/**/*.templ"],
  theme: {
    colors: {
      pkmnbg: "#35374B",
      pkmnfg: "#7DD696",
      pkmnfg2: "#1B2729",
      error: "#ff6188",
      grass: "#9bcc50",
      fire: "#fd7d24",
      water: "#4592c4",
      electric: "#FFD700",
      ice: "#51c4e7",
      poison: "#b97fc9",
      psychic: "#f366b9",
      bug: "#729f3f",
      rock: "#b97fc9",
      ghost: "#7b62a3",
      dragon: "#4592c4",
      dark: "#707070",
      steel: "#9eb7b8",
      fairy: "#fdb9e9",
      fighting: "#d56723",
      flying: "#3dc7ef",
      ground: "#ab9842",
      normal: "#a4acaf",
    },
    fontFamily: {
      pkmn: ["PKMN RBYGSC"],
    },
  },
};
