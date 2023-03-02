import { For, Show } from 'solid-js';
import { Categorie, CategorieItem } from '../types/index';

type SideMenuProps = {
  categories: Categorie[];
  handleActiveCategoriesItems: (items: Categorie[]) => void;
  activeCategorie: Categorie;
  handleItem: (item: CategorieItem) => void;
};

const SideMenu = ({
  categories,
  handleActiveCategoriesItems,
  activeCategorie,
  handleItem,
}: SideMenuProps) => {
  return (
    <section class="flex">
      <aside class="h-full border">
        <ul class="w-52 p-2">
          <For each={categories()}>
            {(item) => (
              <li class="cursor-pointer">
                <input
                  type="radio"
                  checked={false}
                  name="categorie"
                  id={item.categorie}
                  value={item.categorie}
                  onChange={() => {
                    handleActiveCategoriesItems(item);
                  }}
                />
                <label htmlFor={item.categorie}>{item.categorie}</label>
              </li>
            )}
          </For>
        </ul>
      </aside>
      <aside class="h-full border">
        <ul class="w-52 p-2">
          <Show when={activeCategorie} fallback={<p>Please select an item</p>}>
            <For each={activeCategorie().items}>
              {(item) => (
                <li
                  class="p-2 cursor-pointer"
                  onClick={() => {
                    handleItem(item);
                  }}
                >
                  {item.title} ({item.votes})
                </li>
              )}
            </For>
          </Show>
        </ul>
      </aside>
    </section>
  );
};

export default SideMenu;
