<script lang="ts">
  import { SelectRepoRoot, SelectZip, GetImageBase64, ApplyUpdate, GetRepoIndex } from '../wailsjs/go/main/App.js'
  import { onMount } from 'svelte';

  interface ButtonStyle {
    Name: string;
    TextColor: number;
    FillColor: number;
    StrokeColor: number;
    StrokeWidth: number;
    CornerRadius: number;
  }

  interface ControllerLayout {
    Elements: any[];
    ButtonStyles: ButtonStyle[];
    Styles: any;
    Name: string;
    Author: string;
    Description: string;
  }

  interface ParsedPackage {
    ControllerID: string;
    VersionCode: number;
    Layout: ControllerLayout;
    IconPath: string;
    Screenshots: string[];
    IsUpdate: boolean;
    CurrentIndex: IndexEntry | null;
  }

  interface IndexEntry {
    id: string;
    name: string;
    author: string;
    version?: string;
    versionCode?: number;
  }

  let repoRoot = "";
  let pkg: ParsedPackage | null = null;
  let repoIndex: IndexEntry[] = [];
  let iconBase64 = "";
  let screenshotsBase64: string[] = [];

  function intToRGBA(colorInt: number) {
    if (colorInt === 0) return 'transparent';
    const a = ((colorInt >> 24) & 0xff) / 255;
    const r = (colorInt >> 16) & 0xff;
    const g = (colorInt >> 8) & 0xff;
    const b = colorInt & 0xff;
    return `rgba(${r},${g},${b},${a})`;
  }

  function getButtonStyle(styleName: string) {
    if (!pkg || !pkg.Layout || !pkg.Layout.ButtonStyles) return {};
    const style = pkg.Layout.ButtonStyles.find(s => s.Name === styleName);
    if (!style) return {};
    return {
      color: intToRGBA(style.TextColor),
      background: intToRGBA(style.FillColor),
      border: `${style.StrokeWidth}px solid ${intToRGBA(style.StrokeColor)}`,
      borderRadius: `${style.CornerRadius}px`
    };
  }

  async function handleSelectRepo() {
    const res = await SelectRepoRoot();
    if (res) {
      repoRoot = res;
      repoIndex = await GetRepoIndex();
    }
  }

  async function handleSelectZip() {
    const res = await SelectZip();
    if (res) {
      pkg = res as ParsedPackage;
      if (pkg.IconPath) {
        iconBase64 = await GetImageBase64(pkg.IconPath);
      }
      if (pkg.Screenshots) {
        screenshotsBase64 = await Promise.all(
          pkg.Screenshots.map(s => GetImageBase64(s))
        );
      }
    }
  }

  async function handleApply() {
    try {
      await ApplyUpdate();
      alert("Success!");
      repoIndex = await GetRepoIndex();
    } catch (e) {
      alert("Error: " + e);
    }
  }
</script>

<main class="container">
  <div class="sidebar">
    <div class="sidebar-header">
      <h3>控制器列表</h3>
      <button class="btn-small" on:click={handleSelectRepo}>打开仓库</button>
    </div>
    <div class="controller-list">
      {#each repoIndex as entry}
        <div class="controller-item" class:active={pkg?.ControllerID === entry.id}>
          <div class="item-main">
            <span class="name">{entry.name}</span>
            <span class="id">{entry.id}</span>
          </div>
          {#if entry.version}
            <span class="version">v{entry.version}</span>
          {/if}
        </div>
      {/each}
    </div>
    <div class="sidebar-footer">
      <p>{repoRoot || '未选择仓库'}</p>
    </div>
  </div>

  <div class="content">
    <div class="toolbar">
      <button class="btn" on:click={handleSelectZip}>导入 ZIP</button>
      <button class="btn btn-primary" on:click={handleApply} disabled={!pkg}>应用更新</button>
    </div>

    <div class="details">
      {#if pkg}
        <div class="pkg-info">
          <div class="pkg-header">
            <div class="icon-wrap">
              {#if iconBase64}
                <img src="data:image/png;base64,{iconBase64}" alt="icon" class="icon" />
              {/if}
            </div>
            <div class="pkg-title">
              <h2>{pkg.Layout?.Name || pkg.ControllerID}</h2>
              <div class="badges">
                <span class="badge {pkg.IsUpdate ? 'update' : 'new'}">
                  {pkg.IsUpdate ? '更新' : '新增'}
                </span>
              </div>
            </div>
          </div>

          <div class="info-grid">
            <div class="info-item">
              <label>ID</label>
              <span>{pkg.ControllerID}</span>
            </div>
            <div class="info-item">
              <label>作者</label>
              <span>{pkg.Layout?.Author || '未知'}</span>
            </div>
            <div class="info-item">
              <label>版本</label>
              <div class="version-info">
                <span>{pkg.VersionCode}</span>
                {#if pkg.IsUpdate && pkg.CurrentIndex}
                  <span class="version-diff">({pkg.CurrentIndex.versionCode} → {pkg.VersionCode})</span>
                {/if}
              </div>
            </div>
            <div class="info-item full">
              <label>描述</label>
              <span>{pkg.Layout?.Description || '无描述'}</span>
            </div>
          </div>

          <div class="preview-section">
            <h3>布局预览</h3>
            <div class="preview-canvas">
              {#if pkg.Layout && pkg.Layout.ViewGroups}
                {#each pkg.Layout.ViewGroups as group}
                  {#if group.ViewData}
                    {#if group.ViewData.buttonList}
                      {#each group.ViewData.buttonList as btn}
                        <div 
                          class="preview-element"
                          style="
                            left: {btn.baseInfo.xPosition/10}%;
                            top: {btn.baseInfo.yPosition/10}%;
                            width: {btn.baseInfo.percentageWidth?.size/10 || btn.baseInfo.absoluteWidth/10}%;
                            height: {btn.baseInfo.percentageHeight?.size/10 || btn.baseInfo.absoluteHeight/10}%;
                            {Object.entries(getButtonStyle(btn.style)).map(([k, v]) => `${k.replace(/[A-Z]/g, m => '-' + m.toLowerCase())}: ${v}`).join(';')}
                          "
                        >
                          <span class="btn-text">{btn.text}</span>
                        </div>
                      {/each}
                    {/if}
                  {/if}
                {/each}
              {/if}
            </div>
          </div>

            <div class="screenshot-section">
              <h3>截图预览</h3>
              <div class="screenshot-grid">
                {#each screenshotsBase64 as src}
                  <img src="data:image/png;base64,{src}" alt="screenshot" class="screenshot" />
                {/each}
              </div>
            </div>
        </div>
      {:else}
        <div class="empty">
          <p>请选择一个控件 ZIP 包进行预览和审核</p>
          <button class="btn" on:click={handleSelectZip}>导入 ZIP</button>
        </div>
      {/if}
    </div>
  </div>
</main>

<style>
  .container {
    display: flex;
    height: 100vh;
    width: 100vw;
    overflow: hidden;
    background: #1b2636;
    color: white;
  }

  .sidebar {
    width: 260px;
    border-right: 1px solid #334455;
    display: flex;
    flex-direction: column;
    text-align: left;
  }

  .sidebar-header {
    padding: 10px;
    border-bottom: 1px solid #334455;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .controller-list {
    flex: 1;
    overflow-y: auto;
  }

  .controller-item {
    padding: 12px;
    border-bottom: 1px solid #2a3a4a;
    cursor: pointer;
    display: flex;
    justify-content: space-between;
    align-items: center;
    transition: background 0.2s;
  }

  .controller-item:hover {
    background: #2a3a4a;
  }

  .controller-item.active {
    background: #3498db;
    border-bottom-color: transparent;
  }

  .item-main {
    display: flex;
    flex-direction: column;
    gap: 2px;
    text-align: left;
  }

  .controller-item .name {
    font-weight: bold;
    font-size: 14px;
    color: #ffffff;
  }

  .controller-item .id {
    font-size: 11px;
    color: #8899aa;
  }

  .controller-item.active .id {
    color: #e0e0e0;
  }

  .controller-item .version {
    font-size: 11px;
    background: #2a3a4a;
    padding: 2px 6px;
    border-radius: 4px;
    color: #8899aa;
  }

  .controller-item.active .version {
    background: rgba(255, 255, 255, 0.2);
    color: white;
  }

  .sidebar-footer {
    padding: 10px;
    font-size: 11px;
    background: #141d2a;
    word-break: break-all;
  }

  .content {
    flex: 1;
    display: flex;
    flex-direction: column;
  }

  .toolbar {
    padding: 10px;
    background: #242f3f;
    display: flex;
    gap: 10px;
  }

  .details {
    flex: 1;
    overflow-y: auto;
    padding: 20px;
    background: #111a24;
  }

  .pkg-header {
    display: flex;
    gap: 20px;
    align-items: center;
    margin-bottom: 24px;
  }

  .icon-wrap {
    width: 64px;
    height: 64px;
    background: #2a3a4a;
    border-radius: 12px;
    overflow: hidden;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .icon {
    width: 100%;
    height: 100%;
    object-fit: contain;
  }

  .pkg-title h2 {
    margin: 0 0 8px 0;
    font-size: 24px;
  }

  .badge {
    padding: 4px 12px;
    border-radius: 20px;
    font-size: 12px;
    font-weight: bold;
    text-transform: uppercase;
  }

  .badge.new { background: #2ecc71; color: white; }
  .badge.update { background: #3498db; color: white; }

  .info-grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 16px;
    margin-bottom: 30px;
    background: #1b2636;
    padding: 20px;
    border-radius: 12px;
  }

  .info-item {
    display: flex;
    flex-direction: column;
    gap: 4px;
    text-align: left;
  }

  .info-item.full {
    grid-column: span 3;
  }

  .info-item label {
    font-size: 12px;
    color: #8899aa;
    text-transform: uppercase;
  }

  .info-item span {
    font-size: 16px;
    color: #ffffff;
  }

  .version-info {
    display: flex;
    align-items: center;
    gap: 8px;
  }

  .version-diff {
    font-size: 14px;
    color: #3498db;
    font-weight: bold;
  }

  .preview-section {
    margin-bottom: 30px;
    text-align: left;
  }

  .preview-canvas {
    width: 100%;
    aspect-ratio: 16 / 9;
    background: #000;
    position: relative;
    border-radius: 8px;
    overflow: hidden;
    border: 1px solid #334455;
    background-image: 
      linear-gradient(45deg, #111 25%, transparent 25%),
      linear-gradient(-45deg, #111 25%, transparent 25%),
      linear-gradient(45deg, transparent 75%, #111 75%),
      linear-gradient(-45deg, transparent 75%, #111 75%);
    background-size: 20px 20px;
    background-position: 0 0, 0 10px, 10px -10px, -10px 0px;
  }

  .preview-element {
    position: absolute;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 12px;
    overflow: hidden;
    white-space: nowrap;
    box-sizing: border-box;
  }

  .btn-text {
    transform: scale(var(--scale, 1));
  }

  .screenshot-section {
    text-align: left;
  }

  .screenshot-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
    gap: 16px;
    margin-top: 10px;
  }

  .screenshot {
    width: 100%;
    border-radius: 8px;
    border: 1px solid #334455;
    transition: transform 0.2s;
  }

  .screenshot:hover {
    transform: scale(1.05);
    z-index: 10;
  }

  .btn {
    padding: 8px 16px;
    border-radius: 4px;
    border: none;
    background: #334455;
    color: white;
    cursor: pointer;
  }

  .btn-primary {
    background: #007bff;
  }

  .btn:hover {
    opacity: 0.9;
  }

  .btn:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }

  .btn-small {
    padding: 4px 8px;
    font-size: 12px;
    background: #445566;
    border: none;
    color: white;
    border-radius: 2px;
    cursor: pointer;
  }

  .empty {
    height: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    color: #8899aa;
  }
</style>
