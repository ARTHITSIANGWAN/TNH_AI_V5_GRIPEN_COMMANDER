// 🌐 TNH-AI-V5-GRIPEN-COMMANDER: Bridge Interface
export interface Env {
  DB: D1Database;
  TNH_R2: R2Bucket;
  TNH_KV: KVNamespace;
  AI: any;
}

export default {
  // 1. ท่อรับแรงกระแทกจาก HTTP Request (หน้าบ้าน / API)
  async fetch(request: Request, env: Env, ctx: ExecutionContext): Promise<Response> {
    const url = new URL(request.url);

    // ส่งต่อสิทธิ์และข้อมูลไปยังโครงสร้าง Go Logic ที่พอร์ต 2026
    const goEngineUrl = `http://127.0.0.1:2026${url.pathname}${url.search}`;
    
    try {
      console.log(`📡 [Gripen Bridge]: ยิงส่งข้อมูลตัดเลนจราจรตรงเข้า Pure Go Engine -> ${url.pathname}`);
      
      // ดึงพลังจาก AI Binding ของ Cloudflare แปะส่งไปด้วย
      const modifiedRequest = new Request(goEngineUrl, {
        method: request.method,
        headers: {
          ...Object.fromEntries(request.headers),
          "X-D1-Database-ID": "6a8b4373-bf40-4b63-bb02-f612ecbe63b7",
          "X-Gripen-Gate": "PORT_2026_ACTIVE"
        },
        body: request.body
      });

      return await fetch(modifiedRequest);
    } catch (error) {
      return new Response(JSON.stringify({ 
        status: "BRIDGE_ERROR", 
        message: "ท่อสะพานเชื่อมต่อเครื่องยนต์ Go ขัดข้องก้า!",
        error: String(error)
      }), { status: 502, headers: { "Content-Type": "application/json" } });
    }
  },

  // 2. ท่อคุมเวลาสแกนอัตโนมัติทุก 3 นาที (Cron Trigger) ไวระดับ F-16
  async queue(batch: MessageBatch, env: Env, ctx: ExecutionContext): Promise<void> {
    console.log("⏰ [Cron Tick]: ขุนพล JOD ตื่นมาล้างท่อเคลียร์ขยะระบบรอบ 3 นาทีคราบบอส!");
    await fetch("http://127.0.0.1:2026/api/v5/cron-trigger");
  }
};
