/**
 * Generates a random alphanumeric ID.
 *
 * @returns {string} The randomly generated ID.
 */
export function randomId() {
  return Math.random().toString(36).substring(2)
}
