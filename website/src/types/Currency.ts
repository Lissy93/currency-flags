export interface Currency {
  code: string; // ISO 4217 currency code
  name: string; // Currency name
  flag: string; // URL to currency flag SVG
  emoji: string; // Emoji of flag of associated country
  symbol: string; // The currency symbol
}