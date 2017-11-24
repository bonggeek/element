package element

var elementList = map[string]Element{
	"h":  {"H", 1, "Hydrogen", true},
	"he": {"He", 2, "Helium", true},
	"li": {"Li", 3, "Lithium", true},
	"be": {"Be", 4, "Beryllium", true},
	"b":  {"B", 5, "Boron", true},
	"c":  {"C", 6, "Carbon", true},
	"n":  {"N", 7, "Nitrogen", true},
	"o":  {"O", 8, "Oxygen", true},
	"f":  {"F", 9, "Fluorine", true},
	"ne": {"Ne", 10, "Neon", true},
	"na": {"Na", 11, "Sodium", true},
	"mg": {"Mg", 12, "Magnesium", true},
	"al": {"Al", 13, "Aluminum", true},
	"si": {"Si", 14, "Silicon", true},
	"p":  {"P", 15, "Phosphorus", true},
	"s":  {"S", 16, "Sulfur", true},
	"cl": {"Cl", 17, "Chlorine", true},
	"ar": {"Ar", 18, "Argon", true},
	"k":  {"K", 19, "Potassium", true},
	"ca": {"Ca", 20, "Calcium", true},
	"sc": {"Sc", 21, "Scandium", true},
	"ti": {"Ti", 22, "Titanium", true},
	"v":  {"V", 23, "Vanadium", true},
	"cr": {"Cr", 24, "Chromium", true},
	"mn": {"Mn", 25, "Manganese", true},
	"fe": {"Fe", 26, "Iron", true},
	"co": {"Co", 27, "Cobalt", true},
	"ni": {"Ni", 28, "Nickel", true},
	"cu": {"Cu", 29, "Copper", true},
	"zn": {"Zn", 30, "Zinc", true},
	"ga": {"Ga", 31, "Gallium", true},
	"ge": {"Ge", 32, "Germanium", true},
	"as": {"As", 33, "Arsenic", true},
	"se": {"Se", 34, "Selenium", true},
	"br": {"Br", 35, "Bromine", true},
	"kr": {"Kr", 36, "Krypton", true},
	"rb": {"Rb", 37, "Rubidium", true},
	"sr": {"Sr", 38, "Strontium", true},
	"y":  {"Y", 39, "Yttrium", true},
	"zr": {"Zr", 40, "Zirconium", true},
	"nb": {"Nb", 41, "Niobium", true},
	"mo": {"Mo", 42, "Molybdenum", true},
	"tc": {"Tc", 43, "Technetium", true},
	"ru": {"Ru", 44, "Ruthenium", true},
	"rh": {"Rh", 45, "Rhodium", true},
	"pd": {"Pd", 46, "Palladium", true},
	"ag": {"Ag", 47, "Silver", true},
	"cd": {"Cd", 48, "Cadmium", true},
	"in": {"In", 49, "Indium", true},
	"sn": {"Sn", 50, "Tin", true},
	"sb": {"Sb", 51, "Antimony", true},
	"te": {"Te", 52, "Tellurium", true},
	"i":  {"I", 53, "Iodine", true},
	"xe": {"Xe", 54, "Xenon", true},
	"cs": {"Cs", 55, "Cesium", true},
	"ba": {"Ba", 56, "Barium", true},
	"la": {"La", 57, "Lanthanum", true},
	"ce": {"Ce", 58, "Cerium", true},
	"pr": {"Pr", 59, "Praseodymium", true},
	"nd": {"Nd", 60, "Neodymium", true},
	"pm": {"Pm", 61, "Promethium", true},
	"sm": {"Sm", 62, "Samarium", true},
	"eu": {"Eu", 63, "Europium", true},
	"gd": {"Gd", 64, "Gadolinium", true},
	"tb": {"Tb", 65, "Terbium", true},
	"dy": {"Dy", 66, "Dysprosium", true},
	"ho": {"Ho", 67, "Holmium", true},
	"er": {"Er", 68, "Erbium", true},
	"tm": {"Tm", 69, "Thulium", true},
	"yb": {"Yb", 70, "Ytterbium", true},
	"lu": {"Lu", 71, "Lutetium", true},
	"hf": {"Hf", 72, "Hafnium", true},
	"ta": {"Ta", 73, "Tantalum", true},
	"w":  {"W", 74, "Tungsten", true},
	"re": {"Re", 75, "Rhenium", true},
	"os": {"Os", 76, "Osmium", true},
	"ir": {"Ir", 77, "Iridium", true},
	"pt": {"Pt", 78, "Platinum", true},
	"au": {"Au", 79, "Gold", true},
	"hg": {"Hg", 80, "Mercury", true},
	"tl": {"Tl", 81, "Thallium", true},
	"pb": {"Pb", 82, "Lead", true},
	"bi": {"Bi", 83, "Bismuth", true},
	"po": {"Po", 84, "Polonium", true},
	"at": {"At", 85, "Astatine", true},
	"rn": {"Rn", 86, "Radon", true},
	"fr": {"Fr", 87, "Francium", true},
	"ra": {"Ra", 88, "Radium", true},
	"ac": {"Ac", 89, "Actinium", true},
	"th": {"Th", 90, "Thorium", true},
	"pa": {"Pa", 91, "Protactinium", true},
	"u":  {"U", 92, "Uranium", true},
	"np": {"Np", 93, "Neptunium", true},
	"pu": {"Pu", 94, "Plutonium", true},
	"am": {"Am", 95, "Americium", true},
	"cm": {"Cm", 96, "Curium", true},
	"bk": {"Bk", 97, "Berkelium", true},
	"cf": {"Cf", 98, "Californium", true},
	"es": {"Es", 99, "Einsteinium", true},
	"fm": {"Fm", 100, "Fermium", true},
	"md": {"Md", 101, "Mendelevium", true},
	"no": {"No", 102, "Nobelium", true},
	"lr": {"Lr", 103, "Lawrencium", true},
	"rf": {"Rf", 104, "Rutherfordium", true},
	"db": {"Db", 105, "Dubnium", true},
	"sg": {"Sg", 106, "Seaborgium", true},
	"bh": {"Bh", 107, "Bohrium", true},
	"hs": {"Hs", 108, "Hassium", true},
	"mt": {"Mt", 109, "Meitnerium", true},
	"ds": {"Ds", 110, "Darmstadtium", true},
	"rg": {"Rg", 111, "Roentgenium", true},
	"cn": {"Cn", 112, "Copernicium", true},
	"nh": {"Nh", 113, "Nihonium", true},
	"fl": {"Fl", 114, "Flerovium", true},
	"mc": {"Mc", 115, "Moscovium", true},
	"lv": {"Lv", 116, "Livermorium", true},
	"ts": {"Ts", 117, "Tennessine", true},
	"og": {"Og", 118, "Oganesson", true},
	"a": {"A", 119, "Arbitrarium", false},
	"d": {"D", 120, "Doubtium", false},
	"e": {"E", 121, "Exoticium", false},
	"g": {"G", 122, "Grandiosium", false},
	"j": {"J", 123, "Jocerium", false},
	"l": {"L", 124, "Loverium", false},
	"m": {"M", 125, "Magicinium", false},
	"q": {"Q", 126, "Queenium", false},
	"r": {"R", 127, "Royalium", false},
	"t": {"T", 128, "Techium", false},
	"x": {"X", 129, "Misterium", false},
	"z": {"Z", 130, "Zanthium", false},
}
