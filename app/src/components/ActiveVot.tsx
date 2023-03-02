import { CategorieItem } from '../types/index';

type ActiveVotProp = {
  activeVot: CategorieItem;
};

export default function ActiveVot({ activeVot }: ActiveVotProp) {
  return (
    <div>
      <h3 class="font-semibold">{activeVot().title}</h3>
      <p>{activeVot().description}</p>
      <p>Votes: {activeVot().votes}</p>
    </div>
  );
}
