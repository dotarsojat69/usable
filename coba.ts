function firstOccurrence(s: string, x: string): number {
    const lenS = s.length;
    const lenX = x.length;
  
    for (let i = 0; i <= lenS - lenX; i++) {
      let found = true;
  
      for (let j = 0; j < lenX; j++) {
        if (s[i + j] !== x[j]) {
          found = false;
          break;
        }
      }
  
      if (found) {
        return i;
      }
    }
  
    return -1; // Return -1 jika tidak ditemukan
  }
  
  // Contoh penggunaan
  const s = 'juliasamanthantjulia';
  const x = 'ant';
  console.log(firstOccurrence(s, x));  // Output: 8
  