import { createSignal, For, Show } from 'solid-js';
import { Categorie, CategorieItem } from '../types/index';
import { CATEGORIES_MOCK } from '~/mockData';
import SideMenu from '~/components/SideMenu';

export default function Home() {
  const [activeCategoriesItems, setActiveCategoriesItems] = createSignal<
    CategorieItem[]
  >([], {
    equals: false,
  });
  const [activeItem, setActiveItem] = createSignal<CategorieItem | null>(null, {
    equals: false,
  });
  const [categories, setCategories] =
    createSignal<Categorie[]>(CATEGORIES_MOCK);

  const handleActiveCategoriesItems = (items: CategorieItem[]) => {
    setActiveCategoriesItems(items);
  };

  const handleItem = (item: CategorieItem) => {
    setActiveItem(item);
  };

  return (
    <main class="text-gray-70 w-full flex">
      <SideMenu
        categories={categories}
        handleActiveCategoriesItems={handleActiveCategoriesItems}
        activeCategoriesItems={activeCategoriesItems}
        handleItem={handleItem}
      />
      <section class="p-2 w-full border">
        <Show when={activeItem} fallback={<p>Please select an item</p>}>
          <p>{activeItem()?.description}</p>
        </Show>
      </section>
    </main>
  );
}
