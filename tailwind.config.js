/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./internal/templates/**/*.templ"],
  theme: {
    colors: {
      pkmnbg: "#35374B",
      pkmnfg: "#7DD696",
      pkmnfg2: "#1B2729",
      search: "#888DC6",
      searchbtn: "#D9D9D9",

      grass: "#008000",
      fire: "#FF0000",
      water: "#0000FF",
      electric: "#FFD700",
      ice: "#00FFFF",
      poison: "#800080",
      psychic: "#FF1493",
      bug: "#008000",
      rock: "#A52A2A",
      ghost: "#800080",
      dragon: "#0000FF",
      dark: "#000000",
      steel: "#A9A9A9",
      fairy: "#FF69B4",
      fighting: "#FF0000",
      flying: "#0000FF",
      ground: "#A52A2A",
      normal: "#FFFFFF",
    },
    fontFamily: {
      pkmn: ["PKMN RBYGSC"],
    },
  },
};
