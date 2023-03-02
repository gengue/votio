import { For, Show } from 'solid-js';
import { Categorie, CategorieItem } from '../types/index';

type SideMenuProps = {
  categories: Categorie[];
  handleActiveCategoriesItems: (items: CategorieItem[]) => void;
  activeCategoriesItems: CategorieItem[];
  handleItem: (item: CategorieItem) => void;
};

const SideMenu = ({
  categories,
  handleActiveCategoriesItems,
  activeCategoriesItems,
  handleItem,
}: SideMenuProps) => {
  return (
    <section class="flex">
      <aside class="h-full border">
        <ul class="w-52 p-2">
          <For each={categories()}>
            {(item) => (
              <li>
                <input
                  type="checkbox"
                  name={item.categorie}
                  id={item.categorie}
                  onChange={() => {
                    handleActiveCategoriesItems(item.items);
                  }}
                />
                {item.categorie}
              </li>
            )}
          </For>
        </ul>
      </aside>
      <aside class="h-full border">
        <ul class="w-52 p-2">
          <Show
            when={activeCategoriesItems}
            fallback={<p>Please select an item</p>}
          >
            <For each={activeCategoriesItems()}>
              {(item) => (
                <li
                  class="p-2"
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
