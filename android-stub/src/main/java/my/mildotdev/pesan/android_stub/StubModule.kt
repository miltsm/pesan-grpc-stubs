package my.mildotdev.pesan.android_stub

import android.net.Uri
import dagger.Module
import dagger.Provides
import dagger.hilt.InstallIn
import dagger.hilt.components.SingletonComponent
import io.grpc.ManagedChannelBuilder
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.asExecutor
import javax.inject.Qualifier

@Qualifier
@Retention(AnnotationRetention.BINARY)
annotation class PesanService

// This would be ideal, but sadly :app couldn't recognised the provider
//@Module
//@InstallIn(SingletonComponent::class)
//class StubModule {
//    @PesanService
//    @Provides
//    fun providePesanService(): PesanGrpcKt.PesanCoroutineStub {
//        val uri = Uri.parse("http://10.0.2.2:50051")
//        val channel = let {
//            val builder = ManagedChannelBuilder.forAddress(uri.host, uri.port)
//            if (uri.scheme == "https")
//                builder.useTransportSecurity()
//            else
//                builder.usePlaintext()
//           builder.executor(Dispatchers.IO .asExecutor()).build()
//        }
//        return PesanGrpcKt.PesanCoroutineStub(channel)
//    }
//}