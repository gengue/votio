export type Categorie = {
  categorie: string;
  items: CategorieItem[];
};

export type CategorieItem = {
  id: string;
  votes: number;
  title: string;
  description: string;
};
